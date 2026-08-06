package notifications

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler) {
	notificationRoutes := router.Group("/notifications")
	notificationRoutes.Use(middleware.JWTAuthMiddleware())

	{
		// Create notification (mainly for system/admin use)
		notificationRoutes.POST("", handler.CreateNotification)

		// Current user's notifications
		notificationRoutes.GET("", handler.GetUserNotifications)

		// Current user's unread notifications
		notificationRoutes.GET("/unread", handler.GetUnreadNotifications)

		// Single notification
		notificationRoutes.GET("/:id", handler.GetNotification)

		// Mark one notification as read
		notificationRoutes.PUT("/:id/read", handler.MarkAsRead)

		// Mark all notifications as read
		notificationRoutes.PUT("/read-all", handler.MarkAllAsRead)

		// Delete notification
		notificationRoutes.DELETE("/:id", handler.DeleteNotification)
	}
}
