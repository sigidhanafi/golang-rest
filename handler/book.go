package handler

import (
	"golang-rest/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	data := []book.Book{
		{ID: 1, Title: "Psycology of Money"},
		{ID: 2, Title: "Steve Jobs"},
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   data,
	})
}
