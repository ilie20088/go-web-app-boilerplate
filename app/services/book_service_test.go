package services

import (
	"context"
	"testing"

	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/stretchr/testify/assert"
)

var expectedBook = &models.Book{"42", "LotR"}

func TestFetchBook(t *testing.T) {
	service := BookFetcherImpl{}
	InitBookService(&bookRepositoryStub{})

	actualBook, err := service.FetchBook(context.Background(), "1")

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actualBook, expectedBook)
}

type bookRepositoryStub struct{}

func (bookRepositoryStub) FetchBookByID(_ context.Context, _ string) (*models.Book, error) {
	return expectedBook, nil
}
