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
	nrApp, err := utils.InitNewRelic()
	if err != nil {
		log.Fatal(err)
	}
	privateRouter := app.PrivateRouter(chain, nrApp)

	http.Handle("/", privateRouter)
	http.Handle("/_health", pubRouter)

	log.Fatal(http.ListenAndServe(utils.GetAddr(), nil))
}

func init() {
	utils.InitConfig()

	utils.InitLogger()

	services.InitBookService(repositories.BookRepositoryImpl{})
	controllers.InitBooksController(services.BookFetcherImpl{})
}
