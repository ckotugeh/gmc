package video_consultations

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles video consultation HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new video consultation handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateVideoConsultation creates a new video consultation.
func (h *Handler) CreateVideoConsultation(c *gin.Context) {
	var req CreateVideoConsultationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	consultation, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, consultation)
}

// GetVideoConsultation returns a consultation by ID.
func (h *Handler) GetVideoConsultation(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	consultation, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, consultation)
}

// GetDoctorConsultations returns all consultations for a doctor.
func (h *Handler) GetDoctorConsultations(c *gin.Context) {
	doctorID, _ := strconv.ParseUint(c.Param("doctorID"), 10, 64)

	consultations, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, consultations)
}

// GetPatientConsultations returns all consultations for a patient.
func (h *Handler) GetPatientConsultations(c *gin.Context) {
	patientID, _ := strconv.ParseUint(c.Param("patientID"), 10, 64)

	consultations, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, consultations)
}

// GetAllVideoConsultations returns all consultations.
func (h *Handler) GetAllVideoConsultations(c *gin.Context) {
	consultations, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, consultations)
}

// UpdateVideoConsultation updates an existing consultation.
func (h *Handler) UpdateVideoConsultation(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var req UpdateVideoConsultationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	consultation, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, consultation)
}

// DeleteVideoConsultation deletes a consultation.
func (h *Handler) DeleteVideoConsultation(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "video consultation deleted successfully",
	})
}
