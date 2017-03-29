package main

import (
	"log"
	"net/http"

	"github.com/ilie20088/go-web-app-boilerplate/app"
	cfg "github.com/ilie20088/go-web-app-boilerplate/app/config"
)

func main() {
	cfg.SetDefaults()
	err := cfg.ReadFromFile()
	if err != nil {
		log.Fatal(err)
	}
	pubRouter := app.PublicRouter()
	http.Handle("/", pubRouter)

	log.Fatal(http.ListenAndServe(cfg.GetAddr(), pubRouter))
}
