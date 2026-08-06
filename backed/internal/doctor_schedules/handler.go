package doctor_schedules

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles doctor schedule HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new doctor schedule handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateSchedule creates a new doctor schedule.
func (h *Handler) CreateSchedule(c *gin.Context) {
	var req CreateDoctorScheduleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, schedule)
}

// GetSchedule returns a doctor schedule by ID.
func (h *Handler) GetSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule id"})
		return
	}

	schedule, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// GetAllSchedules returns all doctor schedules.
func (h *Handler) GetAllSchedules(c *gin.Context) {
	schedules, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByDoctor returns schedules for a doctor.
func (h *Handler) GetSchedulesByDoctor(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	schedules, err := h.service.GetByDoctorID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// GetSchedulesByDay returns schedules for a weekday.
func (h *Handler) GetSchedulesByDay(c *gin.Context) {
	day := Weekday(c.Param("day"))

	schedules, err := h.service.GetByDay(day)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// UpdateSchedule updates a doctor schedule.
func (h *Handler) UpdateSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule id"})
		return
	}

	var req UpdateDoctorScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schedule, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// DeleteSchedule deletes a doctor schedule.
func (h *Handler) DeleteSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid schedule id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "doctor schedule deleted successfully",
	})
}
