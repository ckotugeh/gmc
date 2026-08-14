package posts

import "github.com/gin-gonic/gin"

func RegisterRoutes(routes *gin.RouterGroup) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	// List all posts
	routes.GET("/posts", handler.GetPosts)

	// Create a post
	routes.POST("/posts", handler.CreatePost)

	// Get a single post
	routes.GET("/posts/:id", handler.GetPost)

	// Get all posts in a community
	routes.GET("/communities/:id/posts", handler.GetCommunityPosts)

	// Update a post
	routes.PUT("/posts/:id", handler.UpdatePost)

	// Delete a post
	routes.DELETE("/posts/:id", handler.DeletePost)
}
