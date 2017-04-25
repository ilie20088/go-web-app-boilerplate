package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"github.com/justinas/alice"
	"github.com/newrelic/go-agent"
)

// PublicRouter creates the application router for public endpoints.
func PublicRouter() *mux.Router {
	r := mux.NewRouter()

	s := r.Path("/_health").Subrouter()
	s.Methods("GET").HandlerFunc(controllers.HealthCheck)

	return r
}

var privateRoutes = []struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}{
	{"GET", "/books/{id:[0-9]+}", controllers.FetchBook},
}

// PrivateRouter creates the application router for private endpoints.
func PrivateRouter(middlewareChain alice.Chain, newRelicApp newrelic.Application) *mux.Router {
	r := mux.NewRouter()

	for _, routeCfg := range privateRoutes {
		var route *mux.Route
		if newRelicApp != nil {
			utils.Logger.Info("Instrumenting " + routeCfg.Method + " " + routeCfg.Path)
			nrPath, nrHandler := newrelic.WrapHandle(newRelicApp, routeCfg.Path, middlewareChain.Then(routeCfg.Handler))
			route = r.Handle(nrPath, nrHandler)
		} else {
			route = r.Handle(routeCfg.Path, middlewareChain.Then(routeCfg.Handler))
		}
		route.Methods(routeCfg.Method)
	}

	return r
}
