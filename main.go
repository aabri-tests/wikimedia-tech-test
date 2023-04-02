package main

import (
	"fmt"
	"log"

	"github.com/wikimedia/cmd/server"
	"github.com/wikimedia/pkg/di"
)

func main() {
	if err := startServer(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
func startServer() error {
	// Build the container using the dependency injection package
	container := di.BuildContainer()

	// Declare a variable to hold the server instance
	var svr *server.Server

	// Invoke the container to get the server instance
	err := container.Invoke(func(s *server.Server) {
		svr = s
	})

	// Check if there was an error invoking the container
	if err != nil {
		return fmt.Errorf("Failed to invoke container: %v", err)
	}

	// Map the routes to the server
	svr.MapRoutes()

	// Start the server
	if err := svr.Start(); err != nil {
		// Return an error if the server failed to start
		return fmt.Errorf("Failed to start server: %v", err)
	}

	// Return nil if the server started successfully
	return nil
}
