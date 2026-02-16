package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"order-service/internal/interceptor"
	pb "order-service/proto"
)

func main() {
	grpcPort := getEnv("GRPC_PORT", "50051")
	metricsPort := getEnv("METRICS_PORT", "9090")

	// --- Create gRPC server with middleware chain ---
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				interceptor.LoggingInterceptor(),
				interceptor.AuthInterceptor(),
				grpc_prometheus.UnaryServerInterceptor,
			),
		),
	)

	// --- Register Prometheus ---
	grpc_prometheus.Register(grpcServer)

	// --- Register your service ---
	orderHandler := NewOrderHandler() // assume constructor exists
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)

	// Enable reflection (useful for grpcurl)
	reflection.Register(grpcServer)

	// --- Start gRPC listener ---
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("gRPC server running on port %s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// --- Metrics HTTP server ---
	metricsServer := &http.Server{
		Addr:    ":" + metricsPort,
		Handler: http.DefaultServeMux,
	}

	http.Handle("/metrics", grpc_prometheus.Handler())

	go func() {
		log.Printf("Metrics server running on port %s", metricsPort)
		if err := metricsServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve metrics: %v", err)
		}
	}()

	// --- Graceful shutdown ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down servers...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	grpcServer.GracefulStop()

	if err := metricsServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("metrics server shutdown failed: %v", err)
	}

	log.Println("Server exited properly")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
