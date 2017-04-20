package app

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	expectedStatusCode := 200
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PrivateRouter()

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	assert.Equal(t, actualStatusCode, expectedStatusCode)
}
