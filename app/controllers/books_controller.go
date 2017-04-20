package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilie20088/go-web-app-boilerplate/app/repositories"
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
)

var bookFetcher services.BookFetcher

// InitBooksController initializes books controller
func InitBooksController(_bookFetcher services.BookFetcher) {
	bookFetcher = _bookFetcher
}

// FetchBook returns a JSON with book or 404 not found
func FetchBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	book, err := bookFetcher.FetchBook(bookID)
	if err == repositories.ErrBookNotFound {
		http.NotFound(w, r)
	} else {
		json.NewEncoder(w).Encode(book)
	}
}
