package lab_requests

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for lab requests.
type Handler struct {
	service Service
}

// NewHandler creates a new handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateLabRequest creates a new lab request.
func (h *Handler) CreateLabRequest(c *gin.Context) {
	var req CreateLabRequestRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, request)
}

// GetLabRequest retrieves a lab request by ID.
func (h *Handler) GetLabRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab request id"})
		return
	}

	request, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

// GetAllLabRequests returns all lab requests.
func (h *Handler) GetAllLabRequests(c *gin.Context) {
	requests, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

// GetLabRequestsByPatient returns lab requests for a patient.
func (h *Handler) GetLabRequestsByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	requests, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

// GetLabRequestsByDoctor returns lab requests created by a doctor.
func (h *Handler) GetLabRequestsByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	requests, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

// GetLabRequestsByAppointment returns lab requests for an appointment.
func (h *Handler) GetLabRequestsByAppointment(c *gin.Context) {
	appointmentID, err := strconv.ParseUint(c.Param("appointment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
		return
	}

	requests, err := h.service.GetByAppointmentID(uint(appointmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

// GetLabRequestsByStatus returns lab requests filtered by status.
func (h *Handler) GetLabRequestsByStatus(c *gin.Context) {
	status := c.Param("status")

	requests, err := h.service.GetByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

// UpdateLabRequest updates a lab request.
func (h *Handler) UpdateLabRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab request id"})
		return
	}

	var req UpdateLabRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

// DeleteLabRequest deletes a lab request.
func (h *Handler) DeleteLabRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab request id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "lab request deleted successfully",
	})
}
