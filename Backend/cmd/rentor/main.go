package main

import (
	"fmt"
	"log"
	"rentor/internal/config"
)

func main() {
	// TODO: logger: zap

	// TODO: config: viper
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config loaded: %+v\n", config)

	// TODO: database: sqlite
	// TODO: router: chi, render
	// TODO: run server
}
