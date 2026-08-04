package appointments

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers appointment routes.
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	appointments := router.Group("/appointments")
	appointments.Use(middleware.JWTAuthMiddleware())
	{
		// Create appointment
		appointments.POST("", handler.CreateAppointment)

		// Get appointments
		appointments.GET("", handler.GetAppointments)
		appointments.GET("/:id", handler.GetAppointment)

		// Doctor appointments
		appointments.GET("/doctor/:doctorID", handler.GetDoctorAppointments)

		// Patient appointments
		appointments.GET("/patient/:patientID", handler.GetPatientAppointments)

		// Update appointment
		appointments.PUT("/:id", handler.UpdateAppointment)

		// Delete appointment
		appointments.DELETE("/:id", handler.DeleteAppointment)
	}
}
