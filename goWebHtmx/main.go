package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")

	handlerOne := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
		io.WriteString(w, r.Method)

	}
	http.HandleFunc("/", handlerOne)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
