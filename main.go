package main

import (
	"booksApi/internal/domain"
	"booksApi/internal/transport"
	"booksApi/pkg/database"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var books []domain.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = database.ConnectDB()
	router := mux.NewRouter()

	transport := transport.Handler{}

	router.HandleFunc("/books", transport.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", transport.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", transport.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", transport.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}", transport.DeleteBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
