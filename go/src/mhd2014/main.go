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
}

// var tasks []*Task

func dailyTasksHandler(rw http.ResponseWriter, req *http.Request) {
	tasks := make([]*Task, 0, 5)

<<<<<<< HEAD
	// tasks = append(tasks, &Task{
	// 	RightImageURL:  "http://www.abendblatt.de/img/vermischtes/crop126086162/8590415521-coriginal/Italian-tourist-died-in-a-bus-accident-in-Gran-Canaria.jpg",
	// 	WrongImageURL:  "http://img.morgenpost.de/img/berlin-aktuell/mobile125855262/6610714992-w1-h1/Matthias-Koeppel-Maler-in-seiner-Galerie-2-.jpg",
	// 	Headline:       "Bus überfährt Touristen - ein Toter und neun Verletzte",
	// 	FullArticleURL: "http://www.abendblatt.de/vermischtes/article126086163/Bus-ueberfaehrt-Touristen-ein-Toter-und-neun-Verletzte.html",
	// 	HintURL:        "http://www.abendblatt.de/vermischtes/article126086163/Bus-ueberfaehrt-Touristen-ein-Toter-und-neun-Verletzte.html",
	// })

	// tasks = append(tasks, &Task{
	// 	RightImageURL:  "http://www.morgenpost.de/vermischtes/stars-und-promis/article126086093/Promi-News-Borchardt-Chef-Mary-modelt-fuer-japanische-Mode.html",
	// 	WrongImageURL:  "http://www.abendblatt.de/img/deutschland/mobile126085746/134071590-w1-h1/Bundestag.jpg",
	// 	Headline:       "Promi-News – \"Borchardt\"-Chef Mary modelt für japanische Mode",
	// 	FullArticleURL: "http://www.morgenpost.de/vermischtes/stars-und-promis/article126086093/Promi-News-Borchardt-Chef-Mary-modelt-fuer-japanische-Mode.html",
	// 	HintURL:        "http://www.morgenpost.de/vermischtes/stars-und-promis/article126086093/Promi-News-Borchardt-Chef-Mary-modelt-fuer-japanische-Mode.html",
	// })

	trendingArticles := GetTrendingArticles(5)
	for _, article := range trendingArticles {
=======
	trendingArticles, tagHistogram := GetTrendingArticles(5)

	for _, article := range trendingArticles {

>>>>>>> FETCH_HEAD
		tasks = append(tasks, &Task{
			RightImageURL:  article.ImageURL,
			WrongImageURL:  article.ImageURL,
			Headline:       article.Headline,
			FullArticleURL: "http://www.morgenpost.de/vermischtes/stars-und-promis/article126086093/Promi-News-Borchardt-Chef-Mary-modelt-fuer-japanische-Mode.html",
			HintURL:        "http://www.morgenpost.de/vermischtes/stars-und-promis/article126086093/Promi-News-Borchardt-Chef-Mary-modelt-fuer-japanische-Mode.html",
		})
	}
<<<<<<< HEAD
=======

	for _, article := range trendingArticles {
		log.Printf("tag scores for article %s  :  \n", article.Headline)
		for _, tag := range article.Tags {
			log.Printf(" %s -> %d \n", tag, tagHistogram[tag])
		}
	}
>>>>>>> FETCH_HEAD

	bytes, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(rw, "%s", bytes)
}
