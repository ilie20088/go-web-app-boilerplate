package controllers

import (
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
	"net/http"
)

var healthService services.HealthService

func Init(hs services.HealthService) {
	healthService = hs
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	healthMsg := healthService.GetHealthMessage()
	w.Write([]byte(healthMsg))
}
