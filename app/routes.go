package app

import (
	"github.com/gorilla/mux"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
)

// PublicRouter creates the application router for public endpoints.
func PublicRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/_health", controllers.HealthCheck)
	return r
}

// PublicRouter creates the application router for public endpoints.
func PrivateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index)
	return r
}
