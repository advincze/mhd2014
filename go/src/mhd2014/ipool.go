package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const isoFormat = "2006-01-02T15:04:05.999Z"

type IPoolArticle struct {
	Id                string
	URL               string
	Category          string
	Tags              []string
	Headline          string
	ImageURL          string
	CreationTimestamp time.Time
	Caption           string
}

func addQuotes(strings []string) []string {
	quotedstrings := make([]string, 0, len(strings))
	for _, s := range strings {

		quotedstrings = append(quotedstrings, `"`+s+`"`)
	}
	return quotedstrings
}

func addNegative(strings []string) []string {
	negativeStr := make([]string, 0, len(strings))
	for _, s := range strings {

		negativeStr = append(negativeStr, "-"+s)
	}
	return negativeStr
}

func SearchIPoolArticles(fromDate, toDate time.Time, publishers []string, limit int, negativeKeywords []string) []*IPoolArticle {

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
		Captions           []string
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
	v.Set("limit", strconv.Itoa(limit))
	if len(publishers) > 0 {
		v.Set("publisher", strings.Join(addQuotes(publishers), ","))
	}
	if len(negativeKeywords) > 0 {
		v.Set("q", `-"`+strings.Join(negativeKeywords, " ")+`"`)
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
			// log.Printf("article externelid already used: %s \n", searchitem.Title)
			continue
		}
		usedExternalIds[searchitem.ExternalIdentifier] = true

		article := &IPoolArticle{
			Id:       searchitem.Id,
			URL:      searchitem.PublishedURL,
			Category: searchitem.Category,
			Caption:  strings.Join(searchitem.Captions, " - "),
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

	return articles
}

func GetTrendingArticles(count int) ([]*IPoolArticle, map[string]int) {

	allArticles := SearchIPoolArticles(time.Now().Add(-2*time.Hour*24), time.Now(), []string{"www.welt.de"}, 500, nil)
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

	return trendingArticles, tagHistogram
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

func getUnrelatedArticle(article *IPoolArticle) *IPoolArticle {
	day := 24 * time.Hour
	t0 := time.Now().Add(-1 * 20 * day)
	unrelArticles := SearchIPoolArticles(t0, t0.Add(10*day), []string{"www.welt.de"}, 1, article.Tags)
	if len(unrelArticles) < 1 {
		return nil
	}
	return unrelArticles[0]
}

func getUnrelatedImageURLs(articles []*IPoolArticle) map[string]string {
	day := 24 * time.Hour
	count := len(articles)
	t0 := time.Now().Add(-1 * time.Duration(20+10*count) * 24 * time.Hour)
	unrelArticles := make(map[string]string, count)
	usedArticlesIds := make(map[string]bool, count)
	for _, article := range articles {
		foundArticles := SearchIPoolArticles(t0, t0.Add(5*day), []string{"www.welt.de", "www.abendblatt.de"}, 20, article.Tags)
		if len(foundArticles) < 1 {
			foundArticles = SearchIPoolArticles(t0, t0.Add(10*day), []string{"www.welt.de", "www.abendblatt.de"}, 20, nil)
		}
		if len(foundArticles) < 1 {
			foundArticles = SearchIPoolArticles(t0, time.Now(), []string{"www.welt.de", "www.abendblatt.de"}, 20, nil)
		}
		if len(foundArticles) > 0 {

			var foundArticle *IPoolArticle
			for i := 0; i < len(foundArticles); i++ {
				if !usedArticlesIds[foundArticles[i].Id] {
					foundArticle = foundArticles[i]
					usedArticlesIds[foundArticle.Id] = true
					break
				}
			}
			unrelArticles[article.Id] = foundArticle.ImageURL
		}
		t0 = t0.Add(10 * day)
	}
	return unrelArticles
}

func getRelatedImageURLs(articles []*IPoolArticle, tagHistogram map[string]int) map[string][]string {
	// imageURLmap := make(map[string][]string, len(articles))

	return nil
}

func getHighestScoringTags(n int, tags []string, tagHistogram map[string]int) []string {
	return nil
}

func toTag(in string) string {
	return strings.ToLower(in)
}
