package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"hello/routes"
)

func main() {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}

	// Create Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	port := "3000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Printf("Starting API server on port %s...\n", port)
	if err := router.Run(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		os.Exit(1)
	}
}
