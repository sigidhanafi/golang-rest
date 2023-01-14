package main

import (
	"golang-rest/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// database connection
	dsn := "root:@tcp(127.0.0.1:3306)/golangweb-gin?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database connection failed")
	}

	log.Println("======Database Connection Success!======")

	// register the router
	router := gin.Default()

	router.GET("/books", handler.GetBooks)

	router.Run()
}
