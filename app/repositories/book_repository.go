package repositories

import (
	"context"
	"errors"

	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"github.com/newrelic/go-agent"
)

// BookRepository interface that defines set of operations supported on book entity
type BookRepository interface {
	FetchBookByID(ctx context.Context, id string) (*models.Book, error)
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
func (BookRepositoryImpl) FetchBookByID(ctx context.Context, id string) (*models.Book, error) {
	defer fetchBookSegment(ctx).End()
	book, ok := storage[id]
	if !ok {
		return nil, ErrBookNotFound
	}

	return book, nil
}

func fetchBookSegment(ctx context.Context) newrelic.DatastoreSegment {
	var segment newrelic.DatastoreSegment
	txn := utils.GetNewRelicTransaction(ctx)
	if txn != nil {
		segment = newrelic.DatastoreSegment{
			StartTime:  newrelic.StartSegmentNow(txn),
			Product:    newrelic.DatastoreMySQL,
			Collection: "books",
			Operation:  "find by id",
		}
	}
	return segment
}
