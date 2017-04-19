package main

import (
	"log"
	"net/http"

	"github.com/ilie20088/go-web-app-boilerplate/app"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"github.com/justinas/alice"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
)

func main() {
	controllers.Init(services.HealthServiceImpl{})
	pubRouter := app.PublicRouter()
	privateRouter := app.PrivateRouter()
	chain := alice.New(app.LoggingMiddleware).Then(privateRouter)

	http.Handle("/", chain)
	http.Handle("/_health", pubRouter)

	log.Fatal(http.ListenAndServe(utils.GetAddr(), nil))
}

func init() {
	// read configurations
	utils.InitConfig()

	// set up logging
	utils.InitLogger()

}
