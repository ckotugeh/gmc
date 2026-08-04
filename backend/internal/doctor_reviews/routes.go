package doctor_reviews

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all doctor review routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	reviews := rg.Group("/doctor-reviews")
	reviews.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		reviews.POST("", handler.CreateDoctorReview)

		// Read
		reviews.GET("", handler.GetAllDoctorReviews)
		reviews.GET("/:id", handler.GetDoctorReview)

		reviews.GET("/doctor/:doctor_id", handler.GetReviewsByDoctor)
		reviews.GET("/doctor/:doctor_id/published", handler.GetPublishedReviewsByDoctor)

		reviews.GET("/patient/:patient_id", handler.GetReviewsByPatient)

		reviews.GET("/appointment/:appointment_id", handler.GetReviewByAppointment)

		// Update
		reviews.PUT("/:id", handler.UpdateDoctorReview)

		// Delete
		reviews.DELETE("/:id", handler.DeleteDoctorReview)
	}
}
