package main

import (
	"golang-rest/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", handler.GetBooks)

	router.Run()
}
