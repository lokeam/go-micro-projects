package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

// API Methods
func getAllFilms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}

func deleteFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range films {

		if item.ID == params["id"] {
			films = append(films[:index], films[index+1:]...)
		}
	}
	// return all existing films after deletion
	json.NewEncoder(w).Encode(films)
}

func getFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range films {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var film Film
	_ = json.NewDecoder(r.Body).Decode(&film)
	film.ID = strconv.Itoa(rand.Intn(10000000000))
	films = append(films, film)
	json.NewEncoder(w).Encode(film)
}

// Note: ideally we won't want to do this when working with databases.
// Given that we are not working with dbs in this example, this achieves
// the functionality that we need
func updateFilm(w http.ResponseWriter, r *http.Request) {
	// set json content type
	w.Header().Set("Content-Type", "application/json")

	// gain access to params
	params := mux.Vars(r)

	// traverse films
	for index, item := range films {
		if item.ID == params["id"] {
			// delete film with sent id
			films = append(films[:index], films[index+1:]...)
			var film Film
			_ = json.NewDecoder(r.Body).Decode(&film)
			film.ID = params["id"]
			// add new film
			films = append(films, film)
			json.NewEncoder(w).Encode(film)
			return
		}
	}
}

func main() {
	router := mux.NewRouter()

	// populate some test films
	films = append(films, Film{ID: "1", Isan: "8943278", Title: "Test Film", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	films = append(films, Film{ID: "2", Isan: "734425", Title: "Test Film, the Second", Director: &Director{Firstname: "Jane", Lastname: "Doh"}})

	// Routes
	router.HandleFunc("/films", getAllFilms).Methods("GET")
	router.HandleFunc("/films/{id}", getFilm).Methods("GET")
	router.HandleFunc("/films", createFilm).Methods("POST")
	router.HandleFunc("/films/{id}", updateFilm).Methods("PUT")
	router.HandleFunc("/films/{id}", deleteFilm).Methods("DELETE")

	fmt.Print("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", router))
}
