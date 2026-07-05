package payments

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all payment routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	payments := rg.Group("/payments")
	payments.Use(middleware.JWTAuthMiddleware())
	{
		// Create
		payments.POST("", handler.CreatePayment)

		// Read
		payments.GET("", handler.GetAllPayments)
		payments.GET("/summary", handler.GetPaymentSummary)
		payments.GET("/:id", handler.GetPayment)

		// Filters
		payments.GET("/appointment/:appointment_id", handler.GetPaymentsByAppointment)
		payments.GET("/patient/:patient_id", handler.GetPaymentsByPatient)
		payments.GET("/doctor/:doctor_id", handler.GetPaymentsByDoctor)
		payments.GET("/hospital/:hospital_id", handler.GetPaymentsByHospital)

		// Update
		payments.PUT("/:id", handler.UpdatePayment)

		// Delete
		payments.DELETE("/:id", handler.DeletePayment)
	}
}
