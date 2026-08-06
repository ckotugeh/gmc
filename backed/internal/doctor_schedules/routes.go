package doctor_schedules

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all doctor schedule routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	schedules := rg.Group("/doctor-schedules")
	schedules.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		schedules.POST("", handler.CreateSchedule)

		// Read
		schedules.GET("", handler.GetAllSchedules)
		schedules.GET("/:id", handler.GetSchedule)

		// Filters
		schedules.GET("/doctor/:doctor_id", handler.GetSchedulesByDoctor)
		schedules.GET("/day/:day", handler.GetSchedulesByDay)

		// Update
		schedules.PUT("/:id", handler.UpdateSchedule)

		// Delete
		schedules.DELETE("/:id", handler.DeleteSchedule)
	}
}
