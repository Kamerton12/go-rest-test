package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testREST/db"
	"testREST/models"
)

func FindBooks(ctx *gin.Context) {
	var books []models.Book
	db.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(ctx *gin.Context) {
	var input models.CreateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}
	db.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(ctx *gin.Context) {
	var book models.Book

	if err := db.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(ctx *gin.Context) {
	var book models.Book
	if err := db.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var inputBook models.UpdateBookInput
	if err := ctx.ShouldBindJSON(&inputBook); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Model(&book).Updates(inputBook)

	ctx.JSON(http.StatusOK,gin.H{"data": book})
}

func DeleteBook(ctx *gin.Context) {
	var book models.Book

	if err := db.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}
