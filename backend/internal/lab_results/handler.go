package lab_results

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for lab results.
type Handler struct {
	service Service
}

// NewHandler creates a new lab result handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateLabResult creates a new lab result.
func (h *Handler) CreateLabResult(c *gin.Context) {
	var req CreateLabResultRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetLabResult retrieves a lab result by ID.
func (h *Handler) GetLabResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab result id"})
		return
	}

	result, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetAllLabResults retrieves all lab results.
func (h *Handler) GetAllLabResults(c *gin.Context) {
	results, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetLabResultsByLabRequest retrieves lab results by lab request ID.
func (h *Handler) GetLabResultsByLabRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("lab_request_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab request id"})
		return
	}

	results, err := h.service.GetByLabRequestID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetLabResultsByPatient retrieves lab results for a patient.
func (h *Handler) GetLabResultsByPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	results, err := h.service.GetByPatientID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetLabResultsByDoctor retrieves lab results by doctor.
func (h *Handler) GetLabResultsByDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	results, err := h.service.GetByDoctorID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// GetLabResultsByStatus retrieves lab results by status.
func (h *Handler) GetLabResultsByStatus(c *gin.Context) {
	status := c.Param("status")

	results, err := h.service.GetByStatus(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}

// UpdateLabResult updates an existing lab result.
func (h *Handler) UpdateLabResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab result id"})
		return
	}

	var req UpdateLabResultRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteLabResult deletes a lab result.
func (h *Handler) DeleteLabResult(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lab result id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "lab result deleted successfully",
	})
}
