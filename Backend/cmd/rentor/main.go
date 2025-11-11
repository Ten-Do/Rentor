package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rentor/internal/config"
	mwLogger "rentor/internal/http-server/middleware"
	"rentor/internal/logger"
	"rentor/internal/storage"

	"github.com/go-chi/chi/v5"
)

func main() {
	// ============================================
	// 1. Configuration loading
	// ============================================
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// ============================================
	// 2. Logger initialization
	// ============================================
	err = logger.InitLogger(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Application started",
		logger.Field("env", cfg.Env),
		logger.Field("storage_path", cfg.StoragePath),
		logger.Field("http_host", cfg.HTTPServer.Host),
		logger.Field("http_port", cfg.HTTPServer.Port),
	)

	// ============================================
	// 3. DB connection
	// ============================================
	db, err := storage.Connect(cfg.StoragePath)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer db.Close()

	logger.Info("Database connected successfully")

	// ============================================
	// 4. Migrations
	// ============================================
	err = storage.RunMigrations(db, "./migrations", "up")
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("Database migrations completed")

	// ============================================
	// 5. Store initialization ( Repositories + Services )
	// ============================================
	// Store includes:
	// - Repositories (working with DB)
	// - Services (business logic)

	// ============================================
	// 6. Routes registration
	// ============================================
	router := chi.NewRouter()

	// ============================================
	// 7. Middlewares registration
	// ============================================
	router.Use(mwLogger.LoggingMiddleware())

	// ============================================
	// 8. HTTP handlers registration
	// ============================================

	// ============================================
	// 9. Server start
	// ============================================

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	server := &http.Server{
		Addr:         cfg.HTTPServer.Host + ":" + cfg.HTTPServer.Port,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.TimeoutSeconds,
		WriteTimeout: cfg.HTTPServer.TimeoutSeconds,
		IdleTimeout:  cfg.HTTPServer.IdleTimeoutSeconds,
	}

	logger.Info("Starting HTTP server", logger.Field("addr", server.Addr))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP server starting failed", logger.Field("error", err.Error()))
		}
	}()

	logger.Info("HTTP server started")

	<-done
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown failed", logger.Field("error", err.Error()))
		return
	}

	logger.Info("Server shutdown completed")
}
