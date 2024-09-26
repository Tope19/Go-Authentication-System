package main

import (
	"go-authentication/config"
	"go-authentication/routes"
	"go-authentication/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	utils.InitLogger()

	r := gin.Default()

	// Initialize database
	config.InitDB()

	// Setup routes
	routes.SetupRoutes(r)

	// Get port from .env file
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8001" // Default port if not specified
	}

	// Run the server
	fmt.Printf("Server running on port %s\n", port)
	r.Run(":" + port)
	utils.Logger.Printf("Server running on port %s\n", port)
    if err := r.Run(":" + port); err != nil {
        utils.Logger.Fatal("Failed to start server:", err)
    }
}