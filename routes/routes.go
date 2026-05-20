package routes

import (
	"github.com/gin-gonic/gin"
	"hello/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes(router *gin.Engine) {
	// User routes
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("", handlers.GetAllUsers)
		userGroup.GET("/:id", handlers.GetUserByID)
		userGroup.POST("", handlers.CreateUser)
		userGroup.PUT("/:id", handlers.UpdateUser)
		userGroup.DELETE("/:id", handlers.DeleteUser)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
}
