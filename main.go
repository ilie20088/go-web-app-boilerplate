package main

import (
	"net/http"
	"github.com/ilie20088/go-web-app-boilerplate/app"
	"log"
)

func main() {
	pubRouter := app.PublicRouter()
	http.Handle("/", pubRouter)

	log.Fatal(http.ListenAndServe(":8000", pubRouter))
}