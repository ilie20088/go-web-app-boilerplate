package main

import (
	"log"
	"net/http"

	"github.com/ilie20088/go-web-app-boilerplate/app"
	"github.com/ilie20088/go-web-app-boilerplate/utils"
)

func main() {
	publicRouterFactory := app.NewPublicRouterFactory()
	pubRouter := publicRouterFactory.PublicRouter()
	privateRouter := app.PrivateRouter()

	http.Handle("/", privateRouter)
	http.Handle("/_health", pubRouter)

	log.Fatal(http.ListenAndServe(utils.GetAddr(), nil))
}

func init() {
	// read configurations
	utils.InitConfig()

	// set up logging
	utils.InitLogger()

}
