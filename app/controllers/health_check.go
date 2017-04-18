package controllers

import (
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
	"net/http"
)

type HealthController struct {
	HealthService services.HealthService
}

func NewHealthController() *HealthController {
	return &HealthController{HealthService: services.NewHealthService()}
}

func (h HealthController) HealthCheck(w http.ResponseWriter, r *http.Request) {
	healthMsg := h.HealthService.GetHealthMessage()
	w.Write([]byte(healthMsg))
}
