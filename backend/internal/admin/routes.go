package admin

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all admin routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	admin := rg.Group("/admin")
	admin.Use(middleware.JWTAuthMiddleware())
	{
		// Dashboard
		admin.GET("/dashboard", handler.GetDashboardStats)

		// Admin audit actions
		admin.POST("", handler.CreateAdminAction)
		admin.GET("", handler.GetAllAdminActions)
		admin.GET("/me", handler.GetMyActions)
		admin.GET("/:id", handler.GetAdminAction)
		admin.PUT("/:id", handler.UpdateAdminAction)
		admin.DELETE("/:id", handler.DeleteAdminAction)
	}
}
