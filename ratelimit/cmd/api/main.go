package main

import (
	"fmt"
	"net/http"
	"ratelimit/config"
	customMiddleware "ratelimit/internal/infra/web/middleware"
	"ratelimit/internal/service"
	"ratelimit/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	limiter := service.NewRedisRateLimiter(
		config.RedisHost,
		config.RedisPort,
		config.RedisDB,
		config.RedisPassword,
		config.RedisRateLimitTTL,
		config.MaxRequestIPSecond,
		config.MaxRequestTokenSecond,
		config.RequestBlockingTimeIP,
		config.RequestBlockingTimeToken,
		config.RateLimiterIPEnabled,
		config.RateLimiterTokenEnabled,
	)

	rateLimiter := customMiddleware.NewRateLimiter(limiter)

	// create web server
	router := chi.NewRouter()

	// use rate limiter middleware
	router.Use(middleware.Logger)
	router.Use(rateLimiter.Middleware)

	// set up routes
	router.Get("/v1/status", usecase.CheckStatus)

	// start web server
	http.ListenAndServe(fmt.Sprintf(":%s", config.WebServerPort), router)

}
