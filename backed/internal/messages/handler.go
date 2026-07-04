package messages

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

// Send message
func (h *Handler) CreateMessage(c *gin.Context) {
	var req CreateMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	senderID := c.GetUint("userID")

	message, err := h.service.CreateMessage(senderID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

// Get single message
func (h *Handler) GetMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	message, err := h.service.GetMessage(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found"})
		return
	}

	c.JSON(http.StatusOK, message)
}

// Get conversation between current user and another user
func (h *Handler) GetConversation(c *gin.Context) {
	otherUserID, _ := strconv.Atoi(c.Param("userID"))
	userID := c.GetUint("userID")

	messages, err := h.service.GetConversation(userID, uint(otherUserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// Get all messages for current user
func (h *Handler) GetUserMessages(c *gin.Context) {
	userID := c.GetUint("userID")

	messages, err := h.service.GetUserMessages(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// Update message
func (h *Handler) UpdateMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req UpdateMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")

	message, err := h.service.UpdateMessage(uint(id), userID, &req)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}

// Delete message
func (h *Handler) DeleteMessage(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	err := h.service.DeleteMessage(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "message deleted successfully"})
}

// Mark message as read
func (h *Handler) MarkAsRead(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	err := h.service.MarkAsRead(uint(id), userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "message marked as read"})
}
