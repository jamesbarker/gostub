package main

import (
	"net/http"
	"text/template"
	"log"
	"strconv"
)

import (
	"flag"
	"net/url"
)

var port = *flag.Int("port", 5001, "Port the stub will listen to")
var templates = template.Must(template.ParseGlob("template/*.json"))

func main() {
	log.Println("Go Stub (-h for options)")
	flag.Parse()
	log.Println("Go Stub Listening on port: " + strconv.Itoa(port))

	http.HandleFunc("/data/2.5/", stubHandler)
	err := http.ListenAndServe(":" + strconv.Itoa(port), nil)
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

	data := map[string]interface{}{
		"cnt": cnt,
	}

	switch(lat) {
	case "-27.468":
		templates.ExecuteTemplate(w, "brisbane", data)
	case "-10.500":
		templates.ExecuteTemplate(w, "badRequest", nil)
	case "-10.404":
		http.NotFound(w, r)
	default:
		templates.ExecuteTemplate(w, "melbourne", data)
	}
}
