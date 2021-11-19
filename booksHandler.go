package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type books struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []books{
	{ID: "1", Title: "The Phoenix Project", Author: "Gene Kim, Kevin Behr, George Spafford"},
	{ID: "2", Title: "Harry Potter and the Prisoner of Azkaban", Author: "J. K. Rowling"},
	{ID: "3", Title: "The Da Vinci Code", Author: "Dan Brown"},
}

func getBooks(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, Books)
}

func booksHandler(c *gin.Context) {
	c.Header("Content-Type", "text/html")
	c.HTML(http.StatusOK, "books.tmpl", gin.H{
		"books": Books,
	})
}
