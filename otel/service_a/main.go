package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"servicea/internal/infra/web"
	"servicea/internal/service"
	"servicea/internal/usecase"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initProvider("service_a", "otel-collector:4317")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	tracer := otel.Tracer("service-tracer")

	port := os.Getenv("APP_PORT")
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Handle("/metrics", promhttp.Handler())
	router.Post("/cep", web.NewValidateCep(
		usecase.NewValidateCep(
			service.NewServiceB(os.Getenv("SERVICE_B_URL"), tracer),
		),
	).Handler)

	go func() {
		if err := http.ListenAndServe(":"+port, router); err != nil {
			log.Fatal(err)
		}
	}()

	select {
	case <-sigCh:
		log.Println("Shutting down gracefully...")
	case <-ctx.Done():
		log.Println("Shutting down...")
	}

	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

}

func initProvider(serviceName, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
		),
	)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, collectorURL, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	tracerExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	bsp := sdktrace.NewBatchSpanProcessor(tracerExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	otel.SetTextMapPropagator(propagation.TraceContext{})

	return tracerProvider.Shutdown, nil

}
