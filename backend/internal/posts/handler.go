package posts

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

// ----------------------------------------------------
// GET /api/posts
// ----------------------------------------------------

func (h *Handler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if posts == nil {
		posts = []Post{}
	}
	c.JSON(http.StatusOK, posts)
}

// ----------------------------------------------------
// POST /api/posts
// ----------------------------------------------------

func (h *Handler) CreatePost(c *gin.Context) {

	var req CreatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	authorID := c.GetUint("userID")

	post, err := h.service.CreatePost(&req, authorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// ----------------------------------------------------
// GET /api/posts/:id
// ----------------------------------------------------

func (h *Handler) GetPost(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post id",
		})
		return
	}

	post, err := h.service.GetPost(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// ----------------------------------------------------
// GET /api/communities/:id/posts
// ----------------------------------------------------

func (h *Handler) GetCommunityPosts(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	posts, err := h.service.GetCommunityPosts(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// ----------------------------------------------------
// PUT /api/posts/:id
// ----------------------------------------------------

func (h *Handler) UpdatePost(c *gin.Context) {

	var req UpdatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post id",
		})
		return
	}

	authorID := c.GetUint("userID")

	post, err := h.service.UpdatePost(uint(id), authorID, &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// ----------------------------------------------------
// DELETE /api/posts/:id
// ----------------------------------------------------

func (h *Handler) DeletePost(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid post id",
		})
		return
	}

	authorID := c.GetUint("userID")

	err = h.service.DeletePost(uint(id), authorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
