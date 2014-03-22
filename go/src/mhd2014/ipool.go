package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const isoFormat = "2006-01-02T15:04:05.999Z"

type IPoolArticle struct {
	Id                string
	Category          string
	Tags              []string
	Headline          string
	ImageURL          string
	CreationTimestamp time.Time
}

func addQuotes(strings []string) []string {
	quotedstrings := make([]string, 0, len(strings))
	for _, s := range strings {

		quotedstrings = append(quotedstrings, `"`+s+`"`)
	}
	return quotedstrings
}

func SearchIPoolArticles(fromDate, toDate time.Time, publishers []string) []*IPoolArticle {

	type Lingo struct {
		Lemma string
	}

	type IPoolSearchItem struct {
		Id                 string
		Category           string
		ExternalIdentifier string
		DateCreated        int64
		PublishedURL       string
		Title              string
		Keywords           []string
		Linguistics        struct {
			Events   []Lingo
			Geos     []Lingo
			Keywords []Lingo
			Orgs     []Lingo
			Persons  []Lingo
		}
		Medias []struct {
			Type       string
			References []struct {
				Url    string
				Width  int
				Height int
			}
		}
	}

	type IPoolSearchResult struct {
		Documents  []IPoolSearchItem
		Pagination struct {
			Total int
		}
	}

	baseURL := "http://ipool-extern.s.asideas.de:9090/api/v2/search"

	v := url.Values{}
	v.Set("startDate", fromDate.Format(isoFormat))
	v.Set("endDate", toDate.Format(isoFormat))
	v.Set("sortBy", "dateCreated")
	v.Set("limit", "500")
	if len(publishers) > 0 {
		v.Set("publisher", strings.Join(addQuotes(publishers), ","))
	}
	queryURL := baseURL + "?" + v.Encode()
	log.Printf("query: %s", queryURL)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	var searchresult IPoolSearchResult

	json.Unmarshal(bytes, &searchresult)

	articles := make([]*IPoolArticle, 0, len(searchresult.Documents))

	usedExternalIds := make(map[string]bool, len(searchresult.Documents))

	log.Printf("search found %d documents, received first %d \n", searchresult.Pagination.Total, len(searchresult.Documents))

	for _, searchitem := range searchresult.Documents {

		if usedExternalIds[searchitem.ExternalIdentifier] {
			log.Printf("article externelid already used: %s \n", searchitem.Title)
			continue
		}
		usedExternalIds[searchitem.ExternalIdentifier] = true

		article := &IPoolArticle{
			Id:       searchitem.Id,
			Category: searchitem.Category,
			// Tags              []Tag
			Headline: searchitem.Title,

			CreationTimestamp: time.Unix(searchitem.DateCreated/1000, 0),
		}

		article.Tags = make([]string, 0, 50)

		for _, keyword := range searchitem.Keywords {
			article.Tags = append(article.Tags, toTag(keyword))
		}
		for _, lin := range searchitem.Linguistics.Events {
			article.Tags = append(article.Tags, toTag(lin.Lemma))
		}
		for _, lin := range searchitem.Linguistics.Geos {
			article.Tags = append(article.Tags, toTag(lin.Lemma))
		}
		for _, lin := range searchitem.Linguistics.Keywords {
			article.Tags = append(article.Tags, toTag(lin.Lemma))
		}
		for _, lin := range searchitem.Linguistics.Orgs {
			article.Tags = append(article.Tags, toTag(lin.Lemma))
		}
		for _, lin := range searchitem.Linguistics.Persons {
			article.Tags = append(article.Tags, toTag(lin.Lemma))
		}

		for _, media := range searchitem.Medias {
			if media.Type == "PICTURE" {
				if len(media.References) > 0 {
					article.ImageURL = media.References[0].Url
					break
				}
			}
		}
		if article.ImageURL != "" {
			articles = append(articles, article)
		} else {
			log.Printf("article has no image: %s \n", article.Headline)
		}

	}

	// log.Printf("searchresults:  %#v\n", articles[0])

	return articles
}

func GetTrendingArticles(count int) []*IPoolArticle {

	allArticles := SearchIPoolArticles(time.Now().Add(-2*time.Hour*24), time.Now(), []string{"www.welt.de"})
	log.Printf("search resulted in %d articles", len(allArticles))
	artmap := make(map[string]*IPoolArticle, len(allArticles))
	for _, art := range allArticles {
		artmap[art.Id] = art
	}
	tagHistogram := getTagHistogram(allArticles)
	articleScoring := getArticleScoring(allArticles, tagHistogram)

	trendingArticles := make([]*IPoolArticle, 0, count)
	for i := 0; i < count; i++ {
		var highscoreArtId string
		var highestScore int
		for id, score := range articleScoring {
			if score > highestScore {
				highestScore = score
				highscoreArtId = id
			}
		}
		trendingArticle := artmap[highscoreArtId]
		log.Printf("trending: %#v \n", trendingArticle)
		delete(articleScoring, highscoreArtId)
		trendingArticles = append(trendingArticles, trendingArticle)
	}

	return trendingArticles
}

func getTagHistogram(articles []*IPoolArticle) map[string]int {
	hist := make(map[string]int)
	for _, article := range articles {
		for _, tag := range article.Tags {
			hist[tag] += 1
		}
	}
	return hist
}

func getArticleScoring(allArticles []*IPoolArticle, tagHistogram map[string]int) map[string]int {
	score := make(map[string]int)
	for _, article := range allArticles {
		for _, tag := range article.Tags {
			score[article.Id] += tagHistogram[tag]
		}
		score[article.Id] /= len(article.Tags)
	}
	return score
}

func toTag(in string) string {
	return strings.ToLower(in)
}
