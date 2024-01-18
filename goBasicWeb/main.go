package main

import (
	"fmt"
	"log"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "sorry, that method isn't supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandlerFunc("/greet", greetHandler)

	fmt.Println("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil (
		log.Fatal(err)
	)
}
