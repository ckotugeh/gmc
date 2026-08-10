package prescriptions

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles prescription HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new prescription handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreatePrescription creates a new prescription.
func (h *Handler) CreatePrescription(c *gin.Context) {
	var req CreatePrescriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prescription, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, prescription)
}

// GetPrescription returns a prescription by ID.
func (h *Handler) GetPrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid prescription id"})
		return
	}

	prescription, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescription)
}

// GetAllPrescriptions returns all prescriptions.
func (h *Handler) GetAllPrescriptions(c *gin.Context) {
	prescriptions, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescriptions)
}

// GetPrescriptionsByDoctor returns prescriptions for a doctor.
func (h *Handler) GetPrescriptionsByDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	prescriptions, err := h.service.GetByDoctorID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescriptions)
}

// GetPrescriptionsByPatient returns prescriptions for a patient.
func (h *Handler) GetPrescriptionsByPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	prescriptions, err := h.service.GetByPatientID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescriptions)
}

// GetPrescriptionsByAppointment returns prescriptions for an appointment.
func (h *Handler) GetPrescriptionsByAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("appointment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
		return
	}

	prescriptions, err := h.service.GetByAppointmentID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescriptions)
}

// UpdatePrescription updates a prescription.
func (h *Handler) UpdatePrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid prescription id"})
		return
	}

	var req UpdatePrescriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prescription, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, prescription)
}

// DeletePrescription deletes a prescription.
func (h *Handler) DeletePrescription(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid prescription id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "prescription deleted successfully",
	})
}
