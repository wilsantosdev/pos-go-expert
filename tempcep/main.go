package main

import (
	"net/http"
	"tempcep/internal/service"
	"tempcep/internal/usecase"
	"tempcep/internal/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/cep/{cep}", web.NewCepTempController(usecase.NewCepTemp(service.NewViaCep(), service.NewWeatherApi("0f169cdbb92a4202838134930241902"))).Handler)

	http.ListenAndServe(":8080", router)

}
