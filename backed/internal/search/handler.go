package search

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles search HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new search handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateSearch records a search and returns its details.
func (h *Handler) CreateSearch(c *gin.Context) {
	var req SearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := c.MustGet("user_id").(uint)

	search, err := h.service.Create(req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, search)
}

// GetSearch returns a search history record by ID.
func (h *Handler) GetSearch(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	search, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, search)
}

// GetMySearches returns the authenticated user's search history.
func (h *Handler) GetMySearches(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	searches, err := h.service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, searches)
}

// GlobalSearch performs a global search.
func (h *Handler) GlobalSearch(c *gin.Context) {
	query := c.Query("q")

	results, err := h.service.GlobalSearch(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, GlobalSearchResponse{
		Query:   query,
		Count:   len(results),
		Results: results,
	})
}

// SearchDoctors searches doctors.
func (h *Handler) SearchDoctors(c *gin.Context) {
	query := c.Query("q")

	results, err := h.service.SearchDoctors(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// SearchPatients searches patients.
func (h *Handler) SearchPatients(c *gin.Context) {
	query := c.Query("q")

	results, err := h.service.SearchPatients(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// SearchHospitals searches hospitals.
func (h *Handler) SearchHospitals(c *gin.Context) {
	query := c.Query("q")

	results, err := h.service.SearchHospitals(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// SearchCommunities searches communities.
func (h *Handler) SearchCommunities(c *gin.Context) {
	query := c.Query("q")

	results, err := h.service.SearchCommunities(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// SearchPosts searches posts.
func (h *Handler) SearchPosts(c *gin.Context) {
	query := c.Query("q")

	results, err := h.service.SearchPosts(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, results)
}

// DeleteSearch deletes a search history record.
func (h *Handler) DeleteSearch(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "search deleted successfully",
	})
}
