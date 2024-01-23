package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Album struct {
	Title  string
	Artist string
}

func main() {
	fmt.Println("Hello World")

	templateHandler := func(w http.ResponseWriter, r *http.Request) {
		webtemplate := template.Must(template.ParseFiles("index.html"))
		/* Sample Album data */
		albums := map[string][]Album{
			"Albums": {
				{Title: "Led Zeppelin", Artist: "Led Zeppelin IV"},
				{Title: "Rolling Stones", Artist: "Tattoo You"},
				{Title: "David Bowie", Artist: "Diamond Dogs"},
				{Title: "The Meters", Artist: "Look-Ky Py Py"},
				{Title: "Journey", Artist: "Escape"},
				{Title: "Otis Redding", Artist: "Otis Blue"},
				{Title: "The Black Crowes", Artist: "The Southern Harmony and Musical Companion"},
			},
		}
		webtemplate.Execute(w, albums)
	}

	albumHandler := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		artist := r.PostFormValue("artist")

		webtemplate := template.Must(template.ParseFiles("index.html"))
		webtemplate.ExecuteTemplate(w, "album-list-element", Album{Title: title, Artist: artist})

	}
	/* Route Handlers */
	http.HandleFunc("/", templateHandler)
	http.HandleFunc("/add-album/", albumHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
