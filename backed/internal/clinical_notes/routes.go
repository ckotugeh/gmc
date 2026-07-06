package clinical_notes

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all clinical note routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	notes := rg.Group("/clinical-notes")
	notes.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		notes.POST("", handler.CreateClinicalNote)

		// Read
		notes.GET("", handler.GetAllClinicalNotes)
		notes.GET("/:id", handler.GetClinicalNote)

		notes.GET("/appointment/:appointment_id", handler.GetClinicalNotesByAppointment)
		notes.GET("/doctor/:doctor_id", handler.GetClinicalNotesByDoctor)
		notes.GET("/patient/:patient_id", handler.GetClinicalNotesByPatient)
		notes.GET("/diagnosis/:diagnosis_id", handler.GetClinicalNotesByDiagnosis)
		notes.GET("/confidential", handler.GetConfidentialClinicalNotes)

		// Update
		notes.PUT("/:id", handler.UpdateClinicalNote)

		// Delete
		notes.DELETE("/:id", handler.DeleteClinicalNote)
	}
}
