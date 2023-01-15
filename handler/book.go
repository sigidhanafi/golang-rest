package handler

import (
	"errors"
	"fmt"
	"golang-rest/book"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (h *bookHandler) DeleteByID(c *gin.Context) {
	ID := c.Param("id")
	idInt, err := strconv.Atoi(ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
	}

	err = h.service.DeleteByID(idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Book deleted",
	})
}

func (h *bookHandler) Create(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		log.Println(err)

		var validationError validator.ValidationErrors
		errorMessage := []string{}

		if errors.As(err, &validationError) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage = append(errorMessage, fmt.Sprintf("Error on field %s condition: %s", e.Field(), e.ActualTag()))
			}
		} else {
			errorMessage = append(errorMessage, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": errorMessage,
		})

		return
	}

	book, err := h.service.Create(bookRequest)

	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err,
		})

		return
	}

	response := map[string]interface{}{
		"success": "OK",
		"message": "Book created",
		"data":    book,
	}

	c.JSON(http.StatusOK, response)
}
