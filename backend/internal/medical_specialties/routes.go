package medical_specialties

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all medical specialty routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	specialties := rg.Group("/medical-specialties")
	specialties.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		specialties.POST("", handler.CreateMedicalSpecialty)

		// Read
		specialties.GET("", handler.GetAllMedicalSpecialties)
		specialties.GET("/active", handler.GetActiveMedicalSpecialties)
		specialties.GET("/:id", handler.GetMedicalSpecialty)
		specialties.GET("/name/:name", handler.GetMedicalSpecialtyByName)
		specialties.GET("/code/:code", handler.GetMedicalSpecialtyByCode)

		// Update
		specialties.PUT("/:id", handler.UpdateMedicalSpecialty)

		// Delete
		specialties.DELETE("/:id", handler.DeleteMedicalSpecialty)
	}
}
