package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var publicdir = flag.String("public", "public", "dir to serve static files")
var port = flag.String("port", "8080", "http port to serve")

func main() {

	flag.Parse()
	http.HandleFunc("/data/test", handler)
	log.Printf("Listening at port %v\n", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "test 123")
}
