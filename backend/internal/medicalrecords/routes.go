package medicalrecords

import (
	"doctor-platform/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes registers all medical record routes.
func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	records := rg.Group("/medical-records")
	records.Use(middleware.JWTAuthMiddleware())
	{
		records.POST("", handler.CreateMedicalRecord)
		records.GET("", handler.GetMedicalRecords)
		records.GET("/:id", handler.GetMedicalRecord)
		records.GET("/patient/:patientID", handler.GetPatientMedicalRecords)
		records.GET("/doctor/:doctorID", handler.GetDoctorMedicalRecords)
		records.PUT("/:id", handler.UpdateMedicalRecord)
		records.DELETE("/:id", handler.DeleteMedicalRecord)
	}
}
