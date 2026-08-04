package vitals

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all vital routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	vitals := rg.Group("/vitals")
	vitals.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		vitals.POST("", handler.CreateVital)

		// Read
		vitals.GET("", handler.GetAllVitals)
		vitals.GET("/:id", handler.GetVital)

		vitals.GET("/patient/:patient_id", handler.GetVitalsByPatient)
		vitals.GET("/doctor/:doctor_id", handler.GetVitalsByDoctor)
		vitals.GET("/appointment/:appointment_id", handler.GetVitalsByAppointment)

		// Update
		vitals.PUT("/:id", handler.UpdateVital)

		// Delete
		vitals.DELETE("/:id", handler.DeleteVital)
	}
}
