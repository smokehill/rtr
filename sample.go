package main

import (
	// "fmt"
	"log"
	"encoding/json"
	"net/http"
	"./rtr"
)

func main() {
	r := rtr.NewRouter()
	r.SetRoute("GET", "/api/books", listBooks)
	r.SetRoute("GET", "/api/books/([0-9]+)", getBook)
	// r.SetRoute("POST", "/api/books", createBook),
	// r.SetRoute("PUT", "/api/books/([0-9]+)", updateBook),
	// r.SetRoute("DELETE", "/api/books/([0-9]+)", deleteBook),
	// ...

	log.Fatal(http.ListenAndServe(":80", r))
}

var books = []Book{
	Book{"1", "My Invents", "Henry Ford"},
	Book{"2", "My Job", "John D. Rockefeller"},
}

type Book struct {
	Id string
	Title string
	Author string
}

// TODO: rest handlers
// TODO: pars url query params

func listBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := rtr.SplitURL(r.URL.String())

	for _, item := range books {
		if item.Id == params[2] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w, r)
}