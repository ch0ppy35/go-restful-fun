package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	appPort := getEnv("APP_PORT", "8080")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")

	// Index
	router.GET("/", indexHandler)
	// Health check endpoint
	router.GET("/healthz", healthCheck)
	// Books ui demo
	router.GET("/books", booksHandler)
	// Books api (wip)
	router.GET("/api/books", getBooks)
	// Redis demo with hit counter
	router.GET("/requests", requestsCounter)

	// Start it up
	fmt.Println("The app will run on port: " + appPort)
	router.Run("0.0.0.0:" + appPort)
}
