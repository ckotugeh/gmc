package comments

import "github.com/gin-gonic/gin"

func RegisterRoutes(routes *gin.RouterGroup) {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	// Comments on posts
	routes.POST("/posts/:id/comments", handler.CreateComment)

	// Get all comments for a post
	routes.GET("/posts/:id/comments", handler.GetPostComments)

	// Update a comment
	routes.PUT("/comments/:id", handler.UpdateComment)

	// Delete a comment
	routes.DELETE("/comments/:id", handler.DeleteComment)
}
