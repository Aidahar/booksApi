package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Concurrency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Parts", Year: "2014"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == id {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if book.ID == item.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if id == item.ID {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
}
