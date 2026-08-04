package lab_requests

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all lab request routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	labRequests := rg.Group("/lab-requests")
	labRequests.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		labRequests.POST("", handler.CreateLabRequest)

		// Read
		labRequests.GET("", handler.GetAllLabRequests)
		labRequests.GET("/:id", handler.GetLabRequest)

		// Filters
		labRequests.GET("/patient/:patient_id", handler.GetLabRequestsByPatient)
		labRequests.GET("/doctor/:doctor_id", handler.GetLabRequestsByDoctor)
		labRequests.GET("/appointment/:appointment_id", handler.GetLabRequestsByAppointment)
		labRequests.GET("/status/:status", handler.GetLabRequestsByStatus)

		// Update
		labRequests.PUT("/:id", handler.UpdateLabRequest)

		// Delete
		labRequests.DELETE("/:id", handler.DeleteLabRequest)
	}
}
