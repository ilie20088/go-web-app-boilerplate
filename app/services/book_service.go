package services

import (
	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/ilie20088/go-web-app-boilerplate/app/repositories"
)

// BookFetcher is an interfaces that defines set of business operations related to books
type BookFetcher interface {
	FetchBook(id string) (*models.Book, error)
}

// BookFetcherImpl is an implementation of BookFetcher
type BookFetcherImpl struct{}

var bookRepository repositories.BookRepository

// InitBookService initializes book service
func InitBookService(_bookRepository repositories.BookRepository) {
	bookRepository = _bookRepository
}

// FetchBook fetches book by id or returns a error returned by book repository
func (_ *BookFetcherImpl) FetchBook(id string) (*models.Book, error) {
	return bookRepository.FetchBookByID(id)
}
