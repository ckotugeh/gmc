package availability

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all availability routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	availability := rg.Group("/availability")
	availability.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		availability.POST("", handler.CreateAvailability)

		// Read
		availability.GET("", handler.GetAllAvailability)
		availability.GET("/:id", handler.GetAvailability)

		// Filters
		availability.GET("/doctor/:doctor_id", handler.GetAvailabilityByDoctor)
		availability.GET("/schedule/:schedule_id", handler.GetAvailabilityBySchedule)
		availability.GET("/date/:date", handler.GetAvailabilityByDate)
		availability.GET("/doctor/:doctor_id/date/:date", handler.GetAvailabilityByDoctorAndDate)

		// Update
		availability.PUT("/:id", handler.UpdateAvailability)

		// Delete
		availability.DELETE("/:id", handler.DeleteAvailability)
	}
}
