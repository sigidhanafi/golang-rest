package handler

import (
	"golang-rest/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) bookHandler {
	return bookHandler{service}
}

func (h *bookHandler) GetBooks(c *gin.Context) {

	// data := []book.Book{
	// 	{ID: 1, Title: "Psycology of Money"},
	// 	{ID: 2, Title: "Steve Jobs"},
	// }

	books, err := h.service.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "BAD Request",
			"error":  err,
		})
	}

	var bookResponse []book.BookResponse
	for _, b := range books {
		bookData := book.BookResponse{ID: b.ID, Title: b.Title, Price: b.Price, Description: b.Description}

		bookResponse = append(bookResponse, bookData)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   bookResponse,
	})
}
