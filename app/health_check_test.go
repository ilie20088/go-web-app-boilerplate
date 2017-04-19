package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := "UP"
	request, err := http.NewRequest("GET", "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	router := PublicRouter()

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	actualBody := responseWriter.Body.String()
	assert.Equal(t, actualStatusCode, expectedStatusCode)
	assert.Equal(t, actualBody, expectedBody)
}
