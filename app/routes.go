package app

import (
	"github.com/gorilla/mux"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
)

// PublicRouter creates the application router for public endpoints.
func PublicRouter() *mux.Router {
	r := mux.NewRouter()

	s := r.Path("/_health").Subrouter()
	s.Methods("GET").HandlerFunc(controllers.HealthCheck)

	return r
}

// PrivateRouter creates the application router for private endpoints.
func PrivateRouter() *mux.Router {
	r := mux.NewRouter()

	s := r.PathPrefix("/book").Path("/{id:[0-9]+}").Subrouter()
	s.Methods("GET").HandlerFunc(controllers.FetchBook)

	return r
}
