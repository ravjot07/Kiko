package main

import (
	"kiko/routes"
	"kiko/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the MongoDB URI from the environment variable
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is required")
	}

	// Connect to MongoDB
	utils.ConnectDatabase(mongoURI)

	// Set up Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterProblemRoutes(router)
	routes.RegisterEvaluationRoutes(router)

	// Start the server
	router.Run(":8080")
}
