package communities

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

// POST /api/communities
func (h *Handler) CreateCommunity(c *gin.Context) {
	var req CreateCommunityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("userID")

	community, err := h.service.CreateCommunity(&req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, ToCommunityResponse(community))
}

// GET /api/communities
func (h *Handler) GetCommunities(c *gin.Context) {
	communities, err := h.service.GetCommunities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToCommunityResponses(communities))
}

// GET /api/communities/:id
func (h *Handler) GetCommunity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	community, err := h.service.GetCommunity(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToCommunityResponse(community))
}

// PUT /api/communities/:id
func (h *Handler) UpdateCommunity(c *gin.Context) {
	var req UpdateCommunityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	community, err := h.service.UpdateCommunity(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ToCommunityResponse(community))
}

// DELETE /api/communities/:id
func (h *Handler) DeleteCommunity(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	if err := h.service.DeleteCommunity(uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "community deleted successfully",
	})
}
