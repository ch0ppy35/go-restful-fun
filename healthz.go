package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	message := map[string]string{"message": "OK"}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, message)
}
