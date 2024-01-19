package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// All data stored locally using structs
type Film struct {
	ID       string    `json:"id"`
	Isan     string    `json:"isan"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var films []Film

func main() {
	router := mux.NewRouter()

	// populate some test films
	films = append(films, Film{ID: "1", Isan: "8943278", Title: "Test Film", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	films = append(films, Film{ID: "2", Isan: "734425", Title: "Test Film, the Second", Director: &Director{Firstname: "Jane", Lastname: "Doh"}})

	// Routes
	router.HandleFunc("/films", getFilms).Methods("GET")
	router.HandleFunc("/films/{id}", getFilm).Methods("GET")
	router.HandleFunc("/films", createFilm).Methods("POST")
	router.HandleFunc("/films/{id}", updateFilm).Methods("PUT")
	router.HandleFunc("/films/{id}", deleteFilm).Methods("DELETE")

	fmt.Print("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
