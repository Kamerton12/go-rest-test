package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"testREST/controllers"
	"testREST/db"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	db.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run(":8080")
}


