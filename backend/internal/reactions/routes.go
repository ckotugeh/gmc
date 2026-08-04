package reactions

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	reactionRoutes := router.Group("")
	reactionRoutes.Use(middleware.JWTAuthMiddleware())

	{
		// Create a reaction (like/dislike)
		reactionRoutes.POST("/posts/:id/reactions", handler.CreateReaction)

		// Get all reactions for a post
		reactionRoutes.GET("/posts/:id/reactions", handler.GetPostReactions)

		// Update an existing reaction
		reactionRoutes.PUT("/reactions/:id", handler.UpdateReaction)

		// Delete a reaction
		reactionRoutes.DELETE("/reactions/:id", handler.DeleteReaction)
	}
}
