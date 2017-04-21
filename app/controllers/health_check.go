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
	if err := CheckDB(); err != nil {
		health.DB = down
		statusCode = http.StatusInternalServerError
	}
	if err := CheckRedis(); err != nil {
		health.Redis = down
		statusCode = http.StatusInternalServerError
	}
	healthJSON, err := json.Marshal(health)
	w.WriteHeader(statusCode)
	if err == nil {
		w.Write(healthJSON)
	}
}

// CheckDB checks if database is available
var CheckDB = func() error {
	session, err := utils.GetDbConnection()
	if err != nil {
		return err
	}
	_, err = session.Exec("SELECT 1") // session.Ping() always returns OK after first successful ping
	if err != nil {
		return err
	}
	return nil
}

// CheckRedis checks if Redis is available
var CheckRedis = func() error {
	statusCmd := utils.RedisClient.Ping()
	return statusCmd.Err()
}
