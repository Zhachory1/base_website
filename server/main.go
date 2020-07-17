// +build ignore

package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var public_dir = flag.String("public", "public", "Directory base name that holds all static files for the website.")

func homeHandler() func(http.ResponseWriter, *http.Request) {
	t, _ := template.ParseFiles(*public_dir + "/index.html")
	data := struct {
		Title string
		Base  string
	}{
		Title: "My page",
		Base: "static",
	}
	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, data)
	}
}

func main() {
	fs := http.FileServer(http.Dir(*public_dir))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
