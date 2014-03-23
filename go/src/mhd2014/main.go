package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var publicdir = flag.String("public", "public", "dir to serve static files")
var port = flag.String("port", "8080", "http port to serve")

func main() {

	flag.Parse()
	http.HandleFunc("/data/test", handler)
	http.HandleFunc("/data/daily", dailyTasksHandler)
	log.Printf("Listening at port %v\n", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "test 123456 ")
}

type Task struct {
	RightImageURL  string `json:"rightImageURL"`
	WrongImageURL  string `json:"wrongImageURL"`
	Headline       string `json:"headline"`
	FullArticleURL string `json:"fullArticleURL"`
	HintURL        string `json:"hintURL"`
	Category       string `json:"category"`
}

// var tasks []*Task

const taskCount = 25

func dailyTasksHandler(rw http.ResponseWriter, req *http.Request) {
	tasks := make([]*Task, 0, taskCount)

	trendingArticles, tagHistogram := GetTrendingArticles(25)

	unrelArticleImageURLs := getUnrelatedImageURLs(trendingArticles)

	for _, article := range trendingArticles {

		tasks = append(tasks, &Task{
			RightImageURL:  article.ImageURL,
			WrongImageURL:  unrelArticleImageURLs[article.Id],
			Headline:       article.Headline,
			FullArticleURL: article.URL,
			HintURL:        "http://www.morgenpost.de/vermischtes/stars-und-promis/article126086093/Promi-News-Borchardt-Chef-Mary-modelt-fuer-japanische-Mode.html",
			Category:       article.Category,
		})
	}

	for _, article := range trendingArticles {
		log.Printf("tag scores for article %s  :  \n", article.Headline)
		for _, tag := range article.Tags {
			log.Printf(" %s -> %d \n", tag, tagHistogram[tag])
		}
	}

	bytes, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(rw, "%s", bytes)
}
