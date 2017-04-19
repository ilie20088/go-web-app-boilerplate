package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := "OK"
	request, err := http.NewRequest("GET", "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	controllers.Init(HealthServiceStub{})
	router := PublicRouter()

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	actualBody := responseWriter.Body.String()
	assert.Equal(t, actualStatusCode, expectedStatusCode)
	assert.Equal(t, actualBody, expectedBody)
}

type HealthServiceStub struct{}

func (h HealthServiceStub) GetHealthMessage() string {
	return "OK"
}
