package presence

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles HTTP requests for presence.
type Handler struct {
	service *Service
}

// NewHandler creates a new presence handler.
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreatePresence creates a presence record for the authenticated user.
func (h *Handler) CreatePresence(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	presence, err := h.service.CreatePresence(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, presence)
}

// GetMyPresence returns the authenticated user's presence.
func (h *Handler) GetMyPresence(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	presence, err := h.service.GetPresenceByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "presence not found",
		})
		return
	}

	c.JSON(http.StatusOK, presence)
}

// GetUserPresence returns another user's presence.
func (h *Handler) GetUserPresence(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	presence, err := h.service.GetPresenceByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "presence not found",
		})
		return
	}

	c.JSON(http.StatusOK, presence)
}

// GetOnlineUsers returns all online users.
func (h *Handler) GetOnlineUsers(c *gin.Context) {
	users, err := h.service.GetOnlineUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdatePresence updates the authenticated user's presence.
func (h *Handler) UpdatePresence(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	var req UpdatePresenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	presence, err := h.service.UpdatePresence(userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presence)
}

// DeletePresence deletes the authenticated user's presence record.
func (h *Handler) DeletePresence(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	presence, err := h.service.GetPresenceByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "presence not found",
		})
		return
	}

	if err := h.service.DeletePresence(presence.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Presence deleted successfully",
	})
}
