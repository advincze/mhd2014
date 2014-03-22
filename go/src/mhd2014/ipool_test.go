package main

import (
	"log"
	"testing"
	"time"
)

func TestIPoolQuery(t *testing.T) {

	t0 := time.Now().Add(-2 * time.Hour * 24) //time.Date(2014, 3, 21, 0, 0, 0, 0, time.UTC)
	SearchIPoolArticles(t0, t0.Add(5*time.Hour*24), []string{"www.welt.de", "www.abendblatt.de"}, 100, []string{"eertrrrt"})
}

func TestTrendingArticles(t *testing.T) {

<<<<<<< HEAD
	trendingArticles := GetTrendingArticles(5)
=======
	trendingArticles, _ := GetTrendingArticles(5)
>>>>>>> FETCH_HEAD
	log.Printf("trending: %v\n", len(trendingArticles))
}
