package transport

import (
	"booksApi/internal/domain"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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

		rows, err := db.Query("select * from books")
		logFatal(err)

		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFatal(err)

			books = append(books, book)
		}

		json.NewEncoder(w).Encode(books)
	}
}

func (h Handler) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		params := mux.Vars(r)

		rows := db.QueryRow("select * from books where id=$1", params["id"])
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		json.NewEncoder(w).Encode(book)
	}
}

func (h Handler) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		var bookId int

		json.NewDecoder(r.Body).Decode(&book)

		err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) returning id;", book.Title, book.Author, book.Year).Scan(&bookId)
		logFatal(err)

		json.NewEncoder(w).Encode(bookId)
	}
}

func (h Handler) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book domain.Book
		json.NewDecoder(r.Body).Decode(&book)

		result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", book.Title, book.Author, book.Year, book.ID)
		logFatal(err)

		rowsUpdated, err := result.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (h Handler) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		result, err := db.Exec("delete from books where id=$1", params["id"])
		logFatal(err)

		rowsDeleted, err := result.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
