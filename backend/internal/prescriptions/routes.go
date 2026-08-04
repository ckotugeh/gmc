package prescriptions

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all prescription routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	prescriptions := rg.Group("/prescriptions")
	prescriptions.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		prescriptions.POST("", handler.CreatePrescription)

		// Read
		prescriptions.GET("", handler.GetAllPrescriptions)
		prescriptions.GET("/:id", handler.GetPrescription)

		// Filter routes
		prescriptions.GET("/doctor/:doctor_id", handler.GetPrescriptionsByDoctor)
		prescriptions.GET("/patient/:patient_id", handler.GetPrescriptionsByPatient)
		prescriptions.GET("/appointment/:appointment_id", handler.GetPrescriptionsByAppointment)

		// Update
		prescriptions.PUT("/:id", handler.UpdatePrescription)

		// Delete
		prescriptions.DELETE("/:id", handler.DeletePrescription)
	}
}
