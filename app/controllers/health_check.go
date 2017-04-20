package controllers

import "net/http"

// HealthCheck responds with 200 OK and UP message indicating that service is up and running
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("UP"))
}
