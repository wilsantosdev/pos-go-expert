package main

import (
	"net/http"
	"os"
	"tempcep/internal/service"
	"tempcep/internal/usecase"
	"tempcep/internal/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	if weatherAPIKey == "" {
		panic("WEATHER_API_KEY must be set")
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/cep/{cep}", web.NewCepTempController(
		usecase.NewCepTemp(
			service.NewViaCep(),
			service.NewWeatherApi(weatherAPIKey),
		),
	).Handler)

	http.ListenAndServe(":8080", router)

}
