package controllers

import (
	"flashcards-backend/database"
	"flashcards-backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get all flashcards
func GetAllFlashcards(c *gin.Context) {
	var flashcards []models.Flashcard
	database.DB.Find(&flashcards)
	c.JSON(http.StatusOK, flashcards)
}

// Create a new flashcard
func CreateFlashcard(c *gin.Context) {
	var flashcard models.Flashcard
	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&flashcard)
	c.JSON(http.StatusOK, flashcard)
}

// Get flashcard by ID
func GetFlashcardByID(c *gin.Context) {
	id := c.Param("id")
	var flashcard models.Flashcard
	if err := database.DB.First(&flashcard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flashcard not found"})
		return
	}
	c.JSON(http.StatusOK, flashcard)
}

// Update flashcard
func UpdateFlashcard(c *gin.Context) {
	id := c.Param("id")
	var flashcard models.Flashcard
	if err := database.DB.First(&flashcard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Flashcard not found"})
		return
	}

	if err := c.ShouldBindJSON(&flashcard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&flashcard)
	c.JSON(http.StatusOK, flashcard)
}

// Delete flashcard
func DeleteFlashcard(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Flashcard{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Flashcard deleted"})
}
