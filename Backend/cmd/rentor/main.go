package main

import (
	"log"
	"rentor/internal/config"
	"rentor/internal/logger"
)

func main() {

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize logger
	err = logger.InitLogger(config.Env)
	if err != nil {
		log.Fatal(err)
	}

	// Log application start
	logger.Info("Application started",
		logger.Field("env", config.Env),
		logger.Field("storage_path", config.StoragePath),
		logger.Field("http_host", config.HTTPServer.Host),
		logger.Field("http_port", config.HTTPServer.Port),
	)

	// TODO: database: sqlite

	// TODO: router: chi, render
	// TODO: run server
}
