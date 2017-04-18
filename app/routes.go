package app

import (
	"github.com/gorilla/mux"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/justinas/alice"
)

type PublicRouterFactory struct {
	healthController *controllers.HealthController
}

func NewPublicRouterFactory() *PublicRouterFactory {
	return &PublicRouterFactory{healthController: controllers.NewHealthController()}
}

// PublicRouter creates the application router for public endpoints.
func (f PublicRouterFactory) PublicRouter() *mux.Router {
	r := mux.NewRouter()
	chain := alice.New(LoggingMiddleware)

	s := r.Path("/_health").Subrouter()
	s.Methods("GET").Handler(chain.ThenFunc(f.healthController.HealthCheck))

	return r
}

// PrivateRouter creates the application router for private endpoints.
func PrivateRouter() *mux.Router {
	r := mux.NewRouter()
	chain := alice.New(LoggingMiddleware)

	s := r.Path("/").Subrouter()
	s.Methods("GET").Handler(chain.ThenFunc(controllers.Index))

	return r
}
