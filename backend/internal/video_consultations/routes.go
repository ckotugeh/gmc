package video_consultations

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all video consultation routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	videoConsultations := rg.Group("/video-consultations")
	videoConsultations.Use(middleware.JWTAuthMiddleware())
	{
		// Create a video consultation
		videoConsultations.POST("", handler.CreateVideoConsultation)

		// Get all consultations
		videoConsultations.GET("", handler.GetAllVideoConsultations)

		// Get a consultation by ID
		videoConsultations.GET("/:id", handler.GetVideoConsultation)

		// Get consultations by doctor
		videoConsultations.GET("/doctor/:doctorID", handler.GetDoctorConsultations)

		// Get consultations by patient
		videoConsultations.GET("/patient/:patientID", handler.GetPatientConsultations)

		// Update a consultation
		videoConsultations.PUT("/:id", handler.UpdateVideoConsultation)

		// Delete a consultation
		videoConsultations.DELETE("/:id", handler.DeleteVideoConsultation)
	}
}
