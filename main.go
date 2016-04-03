package main

import (
	"net/http"
	"text/template"
	"log"
	"strconv"
	"net/url"
	"time"
)

var templates = template.Must(template.ParseGlob("template/*.json"))

func main() {
	log.Println("Go Stub Listening")

	http.HandleFunc("/data/2.5/", stubHandler)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		panic(err.Error())
	}
}

func stubHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json; charset=utf-8");

	urlString := r.URL.String()
	log.Println("Stub hit with: " + urlString)
	urlParams, _ := url.ParseQuery(urlString)

	lat := urlParams["lat"][0]
	cnt, _ := strconv.Atoi(urlParams["cnt"][0])
	data := map[string]interface{}{"cnt": cnt }

	switch(lat) {
	case "-27.468":
		templates.ExecuteTemplate(w, "brisbane", data)
	case "-33.867":
		time.Sleep(5 * time.Second)
		http.ServeFile(w, r, "data/sydney.json")
	case "-10.500":
		http.ServeFile(w, r, "data/badRequest.json")
	case "-10.404":
		http.NotFound(w, r)
	default:
		http.ServeFile(w, r, "data/melbourne.json")
	}
}
