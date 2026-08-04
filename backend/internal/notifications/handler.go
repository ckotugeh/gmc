package notifications

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateNotification creates a new notification.
func (h *Handler) CreateNotification(c *gin.Context) {
	var req CreateNotificationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	notification, err := h.service.CreateNotification(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

// GetNotification returns a single notification.
func (h *Handler) GetNotification(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	notification, err := h.service.GetNotification(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, notification)
}

// GetUserNotifications returns all notifications for the authenticated user.
func (h *Handler) GetUserNotifications(c *gin.Context) {
	userID := c.GetUint("userID")

	notifications, err := h.service.GetUserNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// GetUnreadNotifications returns unread notifications.
func (h *Handler) GetUnreadNotifications(c *gin.Context) {
	userID := c.GetUint("userID")

	notifications, err := h.service.GetUnreadNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

// MarkAsRead marks a notification as read.
func (h *Handler) MarkAsRead(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	if err := h.service.MarkAsRead(uint(id), userID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Notification marked as read",
	})
}

// MarkAllAsRead marks all notifications for the authenticated user as read.
func (h *Handler) MarkAllAsRead(c *gin.Context) {
	userID := c.GetUint("userID")

	if err := h.service.MarkAllAsRead(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All notifications marked as read",
	})
}

// DeleteNotification deletes a notification.
func (h *Handler) DeleteNotification(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	if err := h.service.DeleteNotification(uint(id), userID); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Notification deleted successfully",
	})
}
