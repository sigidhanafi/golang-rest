package handler

import (
	"golang-rest/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	service book.Service
}

func NewBookHandler(service book.Service) bookHandler {
	return bookHandler{service}
}

func (h *bookHandler) GetBooks(c *gin.Context) {

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

func (h *bookHandler) GetBookByID(c *gin.Context) {

	ID := c.Param("id")
	idInt, err := strconv.Atoi(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
	}

	dataBook, err := h.service.FindByID(idInt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
	}

	bookResponse := book.BookResponse{ID: dataBook.ID, Price: dataBook.Price, Description: dataBook.Description, Title: dataBook.Title}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "",
		"data":    bookResponse,
	})
}
