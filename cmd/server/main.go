package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"credit-card-validator/internal/api/grpc"
	"credit-card-validator/internal/api/rest"
	"credit-card-validator/internal/config"
	"credit-card-validator/internal/middleware"
	"credit-card-validator/internal/service"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	grpcserver "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup logger
	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)

	// Create validator service
	validatorService, err := service.NewValidator(&cfg.Validator, logger)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Setup Echo server
	e := echo.New()
	e.HideBanner = true
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Metrics())

	// Setup REST API
	restHandler := rest.NewHandler(validatorService, logger)
	restHandler.RegisterRoutes(e)

	// Serve static files
	e.Static("/", "web")

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "healthy"})
	})

	// Metrics endpoint
	if cfg.MetricsEnabled {
		e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	}

	// Setup gRPC server
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		logger.Fatalf("Failed to listen on gRPC port: %v", err)
	}

	grpcServer := grpcserver.NewServer()
	grpcHandler := grpc.NewServer(validatorService, logger)
	grpcHandler.RegisterServer(grpcServer)
	reflection.Register(grpcServer)

	// Start servers
	var wg sync.WaitGroup

	// Start gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof("Starting gRPC server on port %d", cfg.GRPCPort)
		if err := grpcServer.Serve(grpcListener); err != nil {
			logger.Errorf("gRPC server error: %v", err)
		}
	}()

	// Start HTTP server
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof("Starting HTTP server on port %d", cfg.Port)
		if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			logger.Errorf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down servers...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown HTTP server
	if err := e.Shutdown(ctx); err != nil {
		logger.Errorf("HTTP server shutdown error: %v", err)
	}

	// Shutdown gRPC server
	grpcServer.GracefulStop()

	wg.Wait()
	logger.Info("Servers stopped")
}
