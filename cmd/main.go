package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"homework-5/internal/config"
	"homework-5/internal/pkg/db"
	"homework-5/internal/pkg/pb"
	"homework-5/internal/pkg/repository/postgresql"
	grpc_server "homework-5/internal/pkg/server"
	"log"
	"net"
	"net/http"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(cfg)
	}
	tp, err := tracerProvider("http://" + cfg.TracerHost + ":14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewDB(ctx, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool().Close()
	go http.ListenAndServe(":9091", promhttp.Handler())

	studentRepo := postgresql.NewStudent(database)
	universityRepo := postgresql.NewUniversity(database)

	server := grpc.NewServer()
	pb.RegisterServiceServer(server, grpc_server.NewImplementation(studentRepo, universityRepo))

	if err = server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}

const (
	service     = "api"
	environment = "development"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
		)),
	)
	return tp, nil
}
