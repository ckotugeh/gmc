package doctor_reviews

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles doctor review HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new doctor review handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateDoctorReview creates a new doctor review.
func (h *Handler) CreateDoctorReview(c *gin.Context) {
	var req CreateDoctorReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// GetDoctorReview returns a doctor review by ID.
func (h *Handler) GetDoctorReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid review id"})
		return
	}

	review, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// GetAllDoctorReviews returns all doctor reviews.
func (h *Handler) GetAllDoctorReviews(c *gin.Context) {
	reviews, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReviewsByDoctor returns all reviews for a doctor.
func (h *Handler) GetReviewsByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	reviews, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetPublishedReviewsByDoctor returns all published reviews for a doctor.
func (h *Handler) GetPublishedReviewsByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	reviews, err := h.service.GetPublishedByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReviewsByPatient returns all reviews written by a patient.
func (h *Handler) GetReviewsByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	reviews, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// GetReviewByAppointment returns a review for a specific appointment.
func (h *Handler) GetReviewByAppointment(c *gin.Context) {
	appointmentID, err := strconv.ParseUint(c.Param("appointment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
		return
	}

	review, err := h.service.GetByAppointmentID(uint(appointmentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// UpdateDoctorReview updates a doctor review.
func (h *Handler) UpdateDoctorReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid review id"})
		return
	}

	var req UpdateDoctorReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// DeleteDoctorReview deletes a doctor review.
func (h *Handler) DeleteDoctorReview(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid review id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "doctor review deleted successfully",
	})
}
