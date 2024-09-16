package controllers

import (
	"strconv"

	"github.com/GabrielVictorP/golang-restAPI.git/database"
	"github.com/GabrielVictorP/golang-restAPI.git/models"
	"github.com/gin-gonic/gin"
)

func ShowBooksAll(c *gin.Context) {
	id := c.Param("id")
	newID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID must be a integer",
		})
		return
	}
	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newID).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot fund book" + err.Error(),
		})
		return
	}
	c.JSON(200, book)

}

func CreateBooks(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON" + err.Error(),
		})
		return
	}
	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book" + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list books" + err.Error(),
		})
		return
	}
}
