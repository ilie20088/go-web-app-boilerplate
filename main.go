package main

import (
	"log"
	"net/http"

	"github.com/ilie20088/go-web-app-boilerplate/app"
	"github.com/ilie20088/go-web-app-boilerplate/app/controllers"
	"github.com/ilie20088/go-web-app-boilerplate/app/repositories"
	"github.com/ilie20088/go-web-app-boilerplate/app/services"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
	"github.com/justinas/alice"
)

func main() {
	pubRouter := app.PublicRouter()
	chain := alice.New(app.LoggingMiddleware)
	newRelicApp, err := utils.InitNewRelic()
	if err != nil {
		log.Fatal(err)
	}
	privateRouter := app.PrivateRouter(chain, newRelicApp)

	http.Handle("/", privateRouter)
	http.Handle("/_health", pubRouter)

	log.Fatal(http.ListenAndServe(utils.GetAddr(), nil))
}

func init() {
	// read configurations
	utils.InitConfig()

	// set up logging
	utils.InitLogger()

	// set up connection with cache
	utils.InitCache()

	services.InitBookService(&repositories.BookRepositoryImpl{})
	controllers.InitBooksController(&services.BookFetcherImpl{})
}
