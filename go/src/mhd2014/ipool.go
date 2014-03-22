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
	Tags              []Tag
	Headline          string
	ImageURL          string
	CreationTimestamp time.Time
}

type Tag struct {
	Id        string
	Value     string
	Synonymes []string
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
		Id           string
		Category     string
		DateCreated  int64
		PublishedURL string
		Title        string
		Linguistics  struct {
			Events   Lingo
			Geos     Lingo
			Keywords Lingo
			Orgs     Lingo
			Persons  Lingo
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
		Documents []IPoolSearchItem
	}

	baseURL := "http://ipool-extern.s.asideas.de:9090/api/v2/search"

	v := url.Values{}
	v.Set("startDate", fromDate.Format(isoFormat))
	v.Set("endDate", toDate.Format(isoFormat))
	v.Set("sortBy", "dateCreated")
	if len(publishers) > 0 {
		v.Set("publisher", strings.Join(addQuotes(publishers), ","))
	}
	queryURL := baseURL + "?" + url.QueryEscape(v.Encode())
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

	for _, searchitem := range searchresult.Documents {

		article := &IPoolArticle{
			Id:       searchitem.Id,
			Category: searchitem.Category,
			// Tags              []Tag
			Headline: searchitem.Title,

			CreationTimestamp: time.Unix(searchitem.DateCreated/1000, 0),
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
		}

	}

	log.Printf("searchresults:  %+v\n", searchresult)

	return articles

}
