package reactions

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

// POST /api/posts/:id/reactions
func (h *Handler) CreateReaction(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post ID",
		})
		return
	}

	userID := c.MustGet("userID").(uint)

	var req CreateReactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	reaction, err := h.service.CreateReaction(uint(postID), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, reaction)
}

// GET /api/posts/:id/reactions
func (h *Handler) GetPostReactions(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post ID",
		})
		return
	}

	reactions, err := h.service.GetPostReactions(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reactions)
}

// PUT /api/reactions/:id
func (h *Handler) UpdateReaction(c *gin.Context) {
	reactionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid reaction ID",
		})
		return
	}

	userID := c.MustGet("userID").(uint)

	var req UpdateReactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	reaction, err := h.service.UpdateReaction(uint(reactionID), userID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reaction)
}

// DELETE /api/reactions/:id
func (h *Handler) DeleteReaction(c *gin.Context) {
	reactionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid reaction ID",
		})
		return
	}

	userID := c.MustGet("userID").(uint)

	if err := h.service.DeleteReaction(uint(reactionID), userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reaction deleted successfully",
	})
}
