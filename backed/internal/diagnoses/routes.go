package diagnoses

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all diagnosis routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	diagnoses := rg.Group("/diagnoses")
	diagnoses.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		diagnoses.POST("", handler.CreateDiagnosis)

		// Read
		diagnoses.GET("", handler.GetAllDiagnoses)
		diagnoses.GET("/:id", handler.GetDiagnosis)

		diagnoses.GET("/appointment/:appointment_id", handler.GetDiagnosesByAppointment)
		diagnoses.GET("/doctor/:doctor_id", handler.GetDiagnosesByDoctor)
		diagnoses.GET("/patient/:patient_id", handler.GetDiagnosesByPatient)
		diagnoses.GET("/status/:status", handler.GetDiagnosesByStatus)

		// Update
		diagnoses.PUT("/:id", handler.UpdateDiagnosis)

		// Delete
		diagnoses.DELETE("/:id", handler.DeleteDiagnosis)
	}
}
