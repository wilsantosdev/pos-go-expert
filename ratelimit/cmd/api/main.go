package main

import (
	"fmt"
	"net/http"
	"ratelimit/config"
	"ratelimit/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	// create web server
	router := chi.NewRouter()

	// use rate limiter middleware
	router.Use(middleware.Logger)

	// set up routes
	router.Get("/v1/status", usecase.CheckStatus)

	// start web server
	http.ListenAndServe(fmt.Sprintf(":%s", config.WebServerPort), router)

}
