package allergies

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles allergy HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new allergy handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateAllergy creates a new allergy record.
func (h *Handler) CreateAllergy(c *gin.Context) {
	var req CreateAllergyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allergy, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, allergy)
}

// GetAllergy retrieves an allergy by ID.
func (h *Handler) GetAllergy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid allergy id"})
		return
	}

	allergy, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergy)
}

// GetAllAllergies returns all allergy records.
func (h *Handler) GetAllAllergies(c *gin.Context) {
	allergies, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergies)
}

// GetAllergiesByPatient returns allergies for a patient.
func (h *Handler) GetAllergiesByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	allergies, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergies)
}

// GetAllergiesByDoctor returns allergies recorded by a doctor.
func (h *Handler) GetAllergiesByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	allergies, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergies)
}

// GetAllergiesBySeverity returns allergies filtered by severity.
func (h *Handler) GetAllergiesBySeverity(c *gin.Context) {
	severity := c.Param("severity")

	allergies, err := h.service.GetBySeverity(severity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergies)
}

// GetActiveAllergies returns all active allergies.
func (h *Handler) GetActiveAllergies(c *gin.Context) {
	allergies, err := h.service.GetActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergies)
}

// UpdateAllergy updates an allergy record.
func (h *Handler) UpdateAllergy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid allergy id"})
		return
	}

	var req UpdateAllergyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	allergy, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, allergy)
}

// DeleteAllergy deletes an allergy record.
func (h *Handler) DeleteAllergy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid allergy id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "allergy record deleted successfully",
	})
}
