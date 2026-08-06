package hospitals

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all hospital routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	hospitals := rg.Group("/hospitals")
	hospitals.Use(middleware.JWTAuthMiddleware())
	{
		hospitals.POST("", handler.CreateHospital)
		hospitals.GET("", handler.GetHospitals)
		hospitals.GET("/:id", handler.GetHospital)
		hospitals.PUT("/:id", handler.UpdateHospital)
		hospitals.DELETE("/:id", handler.DeleteHospital)
	}
}
