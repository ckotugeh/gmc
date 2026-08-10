package allergies

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all allergy routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	allergies := rg.Group("/allergies")
	allergies.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		allergies.POST("", handler.CreateAllergy)

		// Read
		allergies.GET("", handler.GetAllAllergies)
		allergies.GET("/:id", handler.GetAllergy)

		allergies.GET("/patient/:patient_id", handler.GetAllergiesByPatient)
		allergies.GET("/doctor/:doctor_id", handler.GetAllergiesByDoctor)
		allergies.GET("/severity/:severity", handler.GetAllergiesBySeverity)
		allergies.GET("/active", handler.GetActiveAllergies)

		// Update
		allergies.PUT("/:id", handler.UpdateAllergy)

		// Delete
		allergies.DELETE("/:id", handler.DeleteAllergy)
	}
}
