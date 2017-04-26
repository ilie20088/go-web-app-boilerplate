package repositories

import (
	"errors"

	"github.com/ilie20088/go-web-app-boilerplate/app/models"
)

// BookRepository interface that defines set of operations supported on book entity
type BookRepository interface {
	FetchBookByID(id string) (*models.Book, error)
}

// BookRepositoryImpl is an implementation of BookRepository
type BookRepositoryImpl struct{}

var storage map[string]*models.Book

// InitBookRepository initialises book repository
func InitBookRepository(_storage map[string]*models.Book) {
	storage = _storage
}

// ErrBookNotFound happens when book is not found in database
var ErrBookNotFound = errors.New("Book not found")

// FetchBookByID fetches book by given id or returns a not found error
func (*BookRepositoryImpl) FetchBookByID(id string) (*models.Book, error) {
	book, ok := storage[id]
	if !ok {
		return nil, ErrBookNotFound
	}

	return book, nil
}
