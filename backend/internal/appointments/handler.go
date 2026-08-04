package appointments

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles appointment HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new appointment handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateAppointment creates a new appointment.
func (h *Handler) CreateAppointment(c *gin.Context) {
	var req CreateAppointmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	patientID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user",
		})
		return
	}

	appointment, err := h.service.CreateAppointment(patientID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ToResponse(appointment))
}

// GetAppointment returns an appointment by ID.
func (h *Handler) GetAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid appointment ID",
		})
		return
	}

	appointment, err := h.service.GetAppointment(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "appointment not found",
		})
		return
	}

	c.JSON(http.StatusOK, ToResponse(appointment))
}

// GetAppointments returns all appointments.
func (h *Handler) GetAppointments(c *gin.Context) {
	appointments, err := h.service.GetAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToResponseList(appointments))
}

// GetDoctorAppointments returns appointments for a doctor.
func (h *Handler) GetDoctorAppointments(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctorID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid doctor ID",
		})
		return
	}

	appointments, err := h.service.GetDoctorAppointments(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToResponseList(appointments))
}

// GetPatientAppointments returns appointments for a patient.
func (h *Handler) GetPatientAppointments(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patientID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid patient ID",
		})
		return
	}

	appointments, err := h.service.GetPatientAppointments(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToResponseList(appointments))
}

// UpdateAppointment updates an appointment.
func (h *Handler) UpdateAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid appointment ID",
		})
		return
	}

	var req UpdateAppointmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	appointment, err := h.service.UpdateAppointment(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToResponse(appointment))
}

// DeleteAppointment deletes an appointment.
func (h *Handler) DeleteAppointment(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid appointment ID",
		})
		return
	}

	if err := h.service.DeleteAppointment(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "appointment deleted successfully",
	})
}
