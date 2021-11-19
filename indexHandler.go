package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var title string = "Go Go Go!"

func indexHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": title,
	})
}
