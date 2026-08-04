package messages

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	messageRoutes := router.Group("/messages")
	messageRoutes.Use(middleware.JWTAuthMiddleware())

	{
		// Send message
		messageRoutes.POST("", handler.CreateMessage)

		// Get all messages for current user
		messageRoutes.GET("", handler.GetUserMessages)

		// Get single message
		messageRoutes.GET("/:id", handler.GetMessage)

		// Get conversation with another user
		messageRoutes.GET("/conversation/:userID", handler.GetConversation)

		// Update message
		messageRoutes.PUT("/:id", handler.UpdateMessage)

		// Delete message
		messageRoutes.DELETE("/:id", handler.DeleteMessage)

		// Mark message as read
		messageRoutes.PUT("/:id/read", handler.MarkAsRead)
	}
}
