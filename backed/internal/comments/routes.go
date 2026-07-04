package comments

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	commentRoutes := router.Group("/api")
	commentRoutes.Use(middleware.JWTAuthMiddleware())
	{
		// Create a comment on a post
		commentRoutes.POST("/posts/:id/comments", handler.CreateComment)

		// Get all comments for a post
		commentRoutes.GET("/posts/:id/comments", handler.GetPostComments)

		// Update a comment
		commentRoutes.PUT("/comments/:id", handler.UpdateComment)

		// Delete a comment
		commentRoutes.DELETE("/comments/:id", handler.DeleteComment)
	}
}
