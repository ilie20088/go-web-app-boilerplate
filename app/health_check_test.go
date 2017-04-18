package app

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
)

func TestHealthCheck(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := "OK"
	request, err := http.NewRequest("GET", "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseWriter := httptest.NewRecorder()
	publicRouterFactory := &PublicRouterFactory{}
	publicRouterFactory.healthController = &controllers.HealthController{HealthService: HealthServiceStub{}}
	router := publicRouterFactory.PublicRouter()

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
