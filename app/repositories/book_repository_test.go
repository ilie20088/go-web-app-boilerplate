package repositories

import (
	"context"
	"testing"

	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/stretchr/testify/assert"
)

func TestFetchBookByIdNotFound(t *testing.T) {
	InitBookRepository(map[string]*models.Book{})
	repository := BookRepositoryImpl{}

	_, err := repository.FetchBookByID(context.Background(), "non-existent-id")

	if err != ErrBookNotFound {
		t.Fatal("Should have returned error ErrBookNotFound")
	}
}

func TestFetchBookByIdFound(t *testing.T) {
	expectedBook := &models.Book{"1", "Lord of the Rings"}
	InitBookRepository(map[string]*models.Book{"1": expectedBook})
	repository := BookRepositoryImpl{}

	actualBook, err := repository.FetchBookByID(context.Background(), "1")

	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedBook, actualBook)
}
