package availability

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler handles availability HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new availability handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateAvailability creates a new availability slot.
func (h *Handler) CreateAvailability(c *gin.Context) {
	var req CreateAvailabilityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slot, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, slot)
}

// GetAvailability returns a slot by ID.
func (h *Handler) GetAvailability(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid availability id"})
		return
	}

	slot, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slot)
}

// GetAllAvailability returns all availability slots.
func (h *Handler) GetAllAvailability(c *gin.Context) {
	slots, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}

// GetAvailabilityByDoctor returns slots for a doctor.
func (h *Handler) GetAvailabilityByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	slots, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}

// GetAvailabilityBySchedule returns slots for a schedule.
func (h *Handler) GetAvailabilityBySchedule(c *gin.Context) {
	scheduleID, err := strconv.ParseUint(c.Param("schedule_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule id"})
		return
	}

	slots, err := h.service.GetByScheduleID(uint(scheduleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}

// GetAvailabilityByDate returns slots for a specific date.
func (h *Handler) GetAvailabilityByDate(c *gin.Context) {
	date, err := time.Parse("2006-01-02", c.Param("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	slots, err := h.service.GetByDate(date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}

// GetAvailabilityByDoctorAndDate returns slots for a doctor on a specific date.
func (h *Handler) GetAvailabilityByDoctorAndDate(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	date, err := time.Parse("2006-01-02", c.Param("date"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	slots, err := h.service.GetByDoctorAndDate(uint(doctorID), date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slots)
}

// UpdateAvailability updates an availability slot.
func (h *Handler) UpdateAvailability(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid availability id"})
		return
	}

	var req UpdateAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	slot, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, slot)
}

// DeleteAvailability deletes an availability slot.
func (h *Handler) DeleteAvailability(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid availability id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "availability slot deleted successfully",
	})
}
