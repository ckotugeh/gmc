package profile

import (
	"net/http"

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

// POST /api/profile
func (h *Handler) CreateProfile(c *gin.Context) {

	var req CreateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("userID")

	profile, err := h.service.CreateProfile(userID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, profile)
}

// GET /api/profile
func (h *Handler) GetProfile(c *gin.Context) {

	userID := c.GetUint("userID")

	profile, err := h.service.GetProfile(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// PUT /api/profile
func (h *Handler) UpdateProfile(c *gin.Context) {

	var req UpdateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.GetUint("userID")

	profile, err := h.service.UpdateProfile(userID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}
