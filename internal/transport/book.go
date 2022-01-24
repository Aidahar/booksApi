package transport

import (
	"booksApi/internal/domain"
	"booksApi/internal/repository/psql"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
}

var books []domain.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (h Handler) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		books = []domain.Book{}
		bookRepo := psql.BookRepository{}

		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (h Handler) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		params := mux.Vars(r)

		books = []domain.Book{}
		bookRepo := psql.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		book = bookRepo.GetBook(db, book, id)

		json.NewEncoder(w).Encode(book)
	}
}

func (h Handler) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := psql.BookRepository{}
		bookID = bookRepo.AddBook(db, book)

		json.NewEncoder(w).Encode(bookID)
	}
}

func (h Handler) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		json.NewDecoder(r.Body).Decode(&book)

		bookRepo := psql.BookRepository{}
		rowsUpdated := bookRepo.UpdateBook(db, book)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (h Handler) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookRepo := psql.BookRepository{}

		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		rowsDeleted := bookRepo.DeleteBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
