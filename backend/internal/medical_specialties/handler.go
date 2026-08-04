package medical_specialties

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles medical specialty HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new medical specialty handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateMedicalSpecialty creates a new medical specialty.
func (h *Handler) CreateMedicalSpecialty(c *gin.Context) {
	var req CreateMedicalSpecialtyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	specialty, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, specialty)
}

// GetMedicalSpecialty returns a medical specialty by ID.
func (h *Handler) GetMedicalSpecialty(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid specialty id"})
		return
	}

	specialty, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialty)
}

// GetAllMedicalSpecialties returns all medical specialties.
func (h *Handler) GetAllMedicalSpecialties(c *gin.Context) {
	specialties, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialties)
}

// GetActiveMedicalSpecialties returns all active medical specialties.
func (h *Handler) GetActiveMedicalSpecialties(c *gin.Context) {
	specialties, err := h.service.GetActive()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialties)
}

// GetMedicalSpecialtyByName returns a medical specialty by name.
func (h *Handler) GetMedicalSpecialtyByName(c *gin.Context) {
	name := c.Param("name")

	specialty, err := h.service.GetByName(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialty)
}

// GetMedicalSpecialtyByCode returns a medical specialty by code.
func (h *Handler) GetMedicalSpecialtyByCode(c *gin.Context) {
	code := c.Param("code")

	specialty, err := h.service.GetByCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialty)
}

// UpdateMedicalSpecialty updates an existing medical specialty.
func (h *Handler) UpdateMedicalSpecialty(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid specialty id"})
		return
	}

	var req UpdateMedicalSpecialtyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	specialty, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialty)
}

// DeleteMedicalSpecialty deletes a medical specialty.
func (h *Handler) DeleteMedicalSpecialty(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid specialty id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "medical specialty deleted successfully",
	})
}
