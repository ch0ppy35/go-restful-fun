package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type books struct {
	ID     string `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []books{
	{ID: "1", ISBN: "978-0988262591", Title: "The Phoenix Project", Author: "Gene Kim, Kevin Behr, George Spafford"},
	{ID: "2", ISBN: "978-0545582933", Title: "Harry Potter and the Prisoner of Azkaban", Author: "J. K. Rowling"},
	{ID: "3", ISBN: "978-0385504201", Title: "The Da Vinci Code", Author: "Dan Brown"},
	{ID: "4", ISBN: "978-1982149482", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
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
