package uploads

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers upload routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	uploads := rg.Group("/uploads")
	uploads.Use(middleware.JWTAuthMiddleware())
	{
		// Upload a file
		uploads.POST("", handler.UploadFile)

		// Get authenticated user's uploads
		uploads.GET("", handler.GetMyUploads)

		// Get a single upload
		uploads.GET("/:id", handler.GetUpload)

		// Update upload metadata
		uploads.PUT("/:id", handler.UpdateUpload)

		// Delete an upload
		uploads.DELETE("/:id", handler.DeleteUpload)
	}
}
