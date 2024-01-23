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
	fmt.Println("Hello world")

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
	http.HandleFunc("/", templateHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
