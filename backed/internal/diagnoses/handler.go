package diagnoses

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles diagnosis HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new diagnosis handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateDiagnosis creates a new diagnosis.
func (h *Handler) CreateDiagnosis(c *gin.Context) {
	var req CreateDiagnosisRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diagnosis, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, diagnosis)
}

// GetDiagnosis returns a diagnosis by ID.
func (h *Handler) GetDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid diagnosis id"})
		return
	}

	diagnosis, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnosis)
}

// GetAllDiagnoses returns all diagnoses.
func (h *Handler) GetAllDiagnoses(c *gin.Context) {
	diagnoses, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnoses)
}

// GetDiagnosesByAppointment returns diagnoses for an appointment.
func (h *Handler) GetDiagnosesByAppointment(c *gin.Context) {
	appointmentID, err := strconv.ParseUint(c.Param("appointment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
		return
	}

	diagnoses, err := h.service.GetByAppointmentID(uint(appointmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnoses)
}

// GetDiagnosesByDoctor returns diagnoses created by a doctor.
func (h *Handler) GetDiagnosesByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	diagnoses, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnoses)
}

// GetDiagnosesByPatient returns diagnoses for a patient.
func (h *Handler) GetDiagnosesByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	diagnoses, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnoses)
}

// GetDiagnosesByStatus returns diagnoses by status.
func (h *Handler) GetDiagnosesByStatus(c *gin.Context) {
	status := c.Param("status")

	diagnoses, err := h.service.GetByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnoses)
}

// UpdateDiagnosis updates a diagnosis.
func (h *Handler) UpdateDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid diagnosis id"})
		return
	}

	var req UpdateDiagnosisRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	diagnosis, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, diagnosis)
}

// DeleteDiagnosis deletes a diagnosis.
func (h *Handler) DeleteDiagnosis(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid diagnosis id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "diagnosis deleted successfully",
	})
}
