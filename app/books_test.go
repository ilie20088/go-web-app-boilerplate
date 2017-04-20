package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/ilie20088/go-web-app-boilerplate/app/repositories"
	"github.com/stretchr/testify/assert"
)

var expectedBook = models.Book{ID: "1", Title: "LotR"}

func TestBookFound(t *testing.T) {
	expectedStatusCode := 200
	request, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter()
	controllers.InitBooksController(bookServiceStub{})

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	assert.Equal(t, actualStatusCode, expectedStatusCode)
	var actualBook models.Book
	err = json.NewDecoder(responseWriter.Body).Decode(&actualBook)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, actualBook, expectedBook)
}

func TestBookNotFound(t *testing.T) {
	expectedStatusCode := 404
	request, err := http.NewRequest("GET", "/book/non-existent", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter()
	controllers.InitBooksController(bookServiceStub{})

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	assert.Equal(t, actualStatusCode, expectedStatusCode)
}

type bookServiceStub struct{}

func (b bookServiceStub) FetchBook(id string) (*models.Book, error) {
	switch id {
	case "1":
		return &expectedBook, nil
	default:
		return nil, repositories.ErrBookNotFound
	}
}
