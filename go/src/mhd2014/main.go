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

func dailyTasksHandler(rw http.ResponseWriter, req *http.Request) {
	tasks := make([]*Task, 0, 5)

	tasks = append(tasks, &Task{
		RightImageURL: "http://www.google.com",
		WrongImageURL: "http://www.google.com",
	})

	bytes, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(rw, "%s", bytes)
}
