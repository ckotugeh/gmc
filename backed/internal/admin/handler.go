package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles admin HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new admin handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateAdminAction creates a new admin audit entry.
func (h *Handler) CreateAdminAction(c *gin.Context) {
	var req CreateAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	adminID, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}

	admin, err := h.service.Create(req, adminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, admin)
}

// GetAdminAction returns a single admin action.
func (h *Handler) GetAdminAction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	admin, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// GetAllAdminActions returns all admin actions.
func (h *Handler) GetAllAdminActions(c *gin.Context) {
	admins, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admins)
}

// GetMyActions returns actions performed by the authenticated admin.
func (h *Handler) GetMyActions(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	adminID, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id"})
		return
	}

	admins, err := h.service.GetByAdminID(adminID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admins)
}

// UpdateAdminAction updates an admin action.
func (h *Handler) UpdateAdminAction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admin)
}

// DeleteAdminAction deletes an admin action.
func (h *Handler) DeleteAdminAction(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "admin action deleted successfully",
	})
}

// GetDashboardStats returns dashboard statistics.
func (h *Handler) GetDashboardStats(c *gin.Context) {
	stats, err := h.service.GetDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
