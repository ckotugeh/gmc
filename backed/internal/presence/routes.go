package presence

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers presence routes.
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	pres := router.Group("/presence")
	pres.Use(middleware.JWTAuthMiddleware())
	{
		// Create a presence record
		pres.POST("", handler.CreatePresence)

		// Get authenticated user's presence
		pres.GET("/me", handler.GetMyPresence)

		// Get another user's presence
		pres.GET("/:userID", handler.GetUserPresence)

		// List online users
		pres.GET("/online", handler.GetOnlineUsers)

		// Update authenticated user's presence
		pres.PUT("", handler.UpdatePresence)

		// Delete authenticated user's presence
		pres.DELETE("", handler.DeletePresence)
	}
}
