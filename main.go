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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set Gin to release mode in production
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Connect to the database
	mongoURI := os.Getenv("MONGO_URI")
	utils.ConnectDatabase(mongoURI)

	// Register routes
	routes.RegisterProblemRoutes(router)
	routes.RegisterEvaluationRoutes(router)

	// Start the server
	router.Run(":8080")
}
