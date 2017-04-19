package controllers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("UP"))
}
