package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	title := "Go Go Go!"
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": title,
	})
}
