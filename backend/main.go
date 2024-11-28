package main

import (
	"flashcards-backend/database"
	"flashcards-backend/models"
	"flashcards-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect Database
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Flashcard{}, &models.Category{})
	seedData()

	// Setup Routes
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start Server
	router.Run(":8080")
}

func seedData() {
	database.DB.Create(&models.Category{
		Name: "testcategorie",
	})

	database.DB.Create(&models.Flashcard{
		CategoryID:    1,
		Title:       "testTitle",
		Description: "testDesc",
		Answers:     "test / test",
	})
}
