package app

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/h2non/gock"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/ilie20088/go-web-app-boilerplate/app/repositories"
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
	"github.com/justinas/alice"
	"github.com/stretchr/testify/assert"
)

var expectedBook = models.Book{ID: "1", Title: "LotR"}
var chain = alice.New(LoggingMiddleware, AuthMiddleware)
var expectedAuthURL = "http://www.google.com"

func TestBookFound(t *testing.T) {
	defer gock.Off()
	gock.New(expectedAuthURL).Get("/").Reply(http.StatusOK)
	expectedStatusCode := http.StatusOK
	request, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter(chain, nil)
	controllers.InitBooksController(&bookServiceStub{})

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
	defer gock.Off()
	gock.New(expectedAuthURL).Get("/").Reply(http.StatusOK)
	expectedStatusCode := http.StatusNotFound
	request, err := http.NewRequest("GET", "/books/non-existent", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter(chain, nil)
	controllers.InitBooksController(&bookServiceStub{})

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	assert.Equal(t, actualStatusCode, expectedStatusCode)
}

func TestBookUnauthorized(t *testing.T) {
	defer gock.Off()
	gock.New(expectedAuthURL).Get("/").Reply(http.StatusUnauthorized)
	expectedStatusCode := http.StatusUnauthorized
	request, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter(chain, nil)
	controllers.InitBooksController(&bookServiceStub{})

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	assert.Equal(t, actualStatusCode, expectedStatusCode)
}

func BenchmarkFetchBook(b *testing.B) {
	defer gock.Off()
	gock.New(expectedAuthURL).Get("/").Reply(http.StatusOK)
	controllers.InitBooksController(&services.BookFetcherImpl{})
	services.InitBookService(&repositories.BookRepositoryImpl{})
	repositories.InitBookRepository(map[string]*models.Book{"1": {"1", "LotR"}})

	request, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		b.Fatal(err)
	}
	router := PrivateRouter(chain, nil)
	for n := 0; n < b.N; n++ {
		responseWriter := httptest.NewRecorder()
		router.ServeHTTP(responseWriter, request)
	}
}

type bookServiceStub struct{}

func (bookServiceStub) FetchBook(_ context.Context, id string) (*models.Book, error) {
	switch id {
	case "1":
		return &expectedBook, nil
	default:
		return nil, repositories.ErrBookNotFound
	}
}
