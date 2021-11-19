package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	appPort := getEnv("APP_PORT", "8080")
	router := gin.Default()

	// Health check endpoint
	router.GET("/healthz", healthCheck)
	// Books crud (wip)
	router.GET("/api/books", getBooks)
	// Redis demo with hit counter
	router.GET("/requests", requestsCounter)

	// Start it up
	fmt.Println("The app will run on port: " + appPort)
	router.Run("0.0.0.0:" + appPort)
}
