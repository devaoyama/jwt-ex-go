package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type post struct {
	Title string `json:"title"`
	Tag string `json:"tag"`
	URL string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/public", public)

	log.Println(http.ListenAndServe(":8080", r))
}

var public = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "ブログのタイトル",
		Tag: "blog",
		URL: "https://qiita.com",
	}
	json.NewEncoder(w).Encode(post)
})
