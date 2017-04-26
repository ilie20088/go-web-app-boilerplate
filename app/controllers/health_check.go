package controllers

import (
	"encoding/json"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"net/http"
)

// HealthStatus stores health information of service 3rd parties
type HealthStatus struct {
	DB    string
	Redis string
}

const (
	up   = "UP"
	down = "DOWN"
)

// HealthCheck checks health of service 3rd parties
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	health := HealthStatus{
		DB:    up,
		Redis: up,
	}
	statusCode := http.StatusOK
	if err := utils.PingDB(); err != nil {
		health.DB = down
		statusCode = http.StatusInternalServerError
	}
	if err := utils.PingRedis(); err != nil {
		health.Redis = down
		statusCode = http.StatusInternalServerError
	}
	healthJSON, err := json.Marshal(health)
	w.WriteHeader(statusCode)
	if err == nil {
		w.Write(healthJSON)
	}
}
