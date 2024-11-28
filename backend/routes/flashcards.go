package routes

import (
	"flashcards-backend/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	flashcards := router.Group("/api/flashcards")
	{
		flashcards.GET("/", controllers.GetAllFlashcards)
		flashcards.POST("/", controllers.CreateFlashcard)
		flashcards.GET("/:id", controllers.GetFlashcardByID)
		flashcards.PUT("/:id", controllers.UpdateFlashcard)
		flashcards.DELETE("/:id", controllers.DeleteFlashcard)
	}
}
