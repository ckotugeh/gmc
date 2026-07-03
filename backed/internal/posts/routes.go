package posts

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	postRoutes := router.Group("/api")
	postRoutes.Use(middleware.JWTAuthMiddleware())
	{
		// Create a post
		postRoutes.POST("/posts", handler.CreatePost)

		// Get a single post
		postRoutes.GET("/posts/:id", handler.GetPost)

		// Get all posts in a community
		postRoutes.GET("/communities/:id/posts", handler.GetCommunityPosts)

		// Update a post
		postRoutes.PUT("/posts/:id", handler.UpdatePost)

		// Delete a post
		postRoutes.DELETE("/posts/:id", handler.DeletePost)
	}
}
