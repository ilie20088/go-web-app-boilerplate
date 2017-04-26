package app

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckUP(t *testing.T) {
	expectedStatusCode := http.StatusOK
	expectedBody := controllers.HealthStatus{DB: "UP", Redis: "UP"}
	request, err := http.NewRequest("GET", "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}
	utils.PingDB = func() error { return nil }
	utils.PingRedis = func() error { return nil }
	responseWriter := httptest.NewRecorder()
	router := PublicRouter()

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	var actualBody controllers.HealthStatus
	json.NewDecoder(responseWriter.Body).Decode(&actualBody)
	assert.Equal(t, actualStatusCode, expectedStatusCode)
	assert.Equal(t, actualBody, expectedBody)
}

func TestHealthCheckDown(t *testing.T) {
	expectedStatusCode := http.StatusInternalServerError
	expectedBody := controllers.HealthStatus{DB: "DOWN", Redis: "UP"}
	request, err := http.NewRequest("GET", "/_health", nil)
	if err != nil {
		t.Fatal(err)
	}
	utils.PingDB = func() error { return errors.New("DB is down") }
	utils.PingRedis = func() error { return nil }
	responseWriter := httptest.NewRecorder()
	router := PublicRouter()

	router.ServeHTTP(responseWriter, request)

	actualStatusCode := responseWriter.Code
	var actualBody controllers.HealthStatus
	json.NewDecoder(responseWriter.Body).Decode(&actualBody)
	assert.Equal(t, actualStatusCode, expectedStatusCode)
	assert.Equal(t, actualBody, expectedBody)
}
