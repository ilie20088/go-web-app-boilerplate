package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/app/models"
	"github.com/ilie20088/go-web-app-boilerplate/app/repositories"
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
	"github.com/justinas/alice"
	"github.com/stretchr/testify/assert"
)

var expectedBook = models.Book{ID: "1", Title: "LotR"}
var chain = alice.New(LoggingMiddleware)

func TestBookFound(t *testing.T) {
	expectedStatusCode := http.StatusOK
	request, err := http.NewRequest("GET", "/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter(chain, nil)
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
	expectedStatusCode := http.StatusNotFound
	request, err := http.NewRequest("GET", "/books/non-existent", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter(chain, nil)
	controllers.InitBooksController(bookServiceStub{})

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	assert.Equal(t, actualStatusCode, expectedStatusCode)
}

func BenchmarkFetchBook(b *testing.B) {
	controllers.InitBooksController(services.BookFetcherImpl{})
	services.InitBookService(repositories.BookRepositoryImpl{})
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

func (b bookServiceStub) FetchBook(id string) (*models.Book, error) {
	switch id {
	case "1":
		return &expectedBook, nil
	default:
		return nil, repositories.ErrBookNotFound
	}
}
