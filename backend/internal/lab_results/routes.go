package lab_results

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all lab result routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	labResults := rg.Group("/lab-results")
	labResults.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		labResults.POST("", handler.CreateLabResult)

		// Read
		labResults.GET("", handler.GetAllLabResults)
		labResults.GET("/:id", handler.GetLabResult)

		// Filters
		labResults.GET("/lab-request/:lab_request_id", handler.GetLabResultsByLabRequest)
		labResults.GET("/patient/:patient_id", handler.GetLabResultsByPatient)
		labResults.GET("/doctor/:doctor_id", handler.GetLabResultsByDoctor)
		labResults.GET("/status/:status", handler.GetLabResultsByStatus)

		// Update
		labResults.PUT("/:id", handler.UpdateLabResult)

		// Delete
		labResults.DELETE("/:id", handler.DeleteLabResult)
	}
}
