package uploads

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler handles upload HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new upload handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// UploadFile uploads a file.
func (h *Handler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	userIDValue, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID := userIDValue.(uint)

	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory"})
		return
	}

	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	upload := &Upload{
		UserID:       userID,
		FileName:     fileName,
		OriginalName: file.Filename,
		FileType:     file.Header.Get("Content-Type"),
		FileSize:     file.Size,
		FilePath:     filePath,
		Description:  c.PostForm("description"),
	}

	if value := c.PostForm("appointment_id"); value != "" {
		if id, err := strconv.ParseUint(value, 10, 64); err == nil {
			v := uint(id)
			upload.AppointmentID = &v
		}
	}

	if value := c.PostForm("medical_record_id"); value != "" {
		if id, err := strconv.ParseUint(value, 10, 64); err == nil {
			v := uint(id)
			upload.MedicalRecordID = &v
		}
	}

	if value := c.PostForm("hospital_id"); value != "" {
		if id, err := strconv.ParseUint(value, 10, 64); err == nil {
			v := uint(id)
			upload.HospitalID = &v
		}
	}

	if value := c.PostForm("is_public"); value == "true" {
		upload.IsPublic = true
	}

	created, err := h.service.Create(upload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toResponse(created))
}

// GetUpload retrieves an upload by ID.
func (h *Handler) GetUpload(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	upload, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(upload))
}

// GetMyUploads retrieves uploads belonging to the authenticated user.
func (h *Handler) GetMyUploads(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	uploads, err := h.service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]UploadResponse, 0, len(uploads))
	for _, upload := range uploads {
		u := upload
		response = append(response, toResponse(&u))
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUpload updates upload metadata.
func (h *Handler) UpdateUpload(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var req UpdateUploadRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upload, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toResponse(upload))
}

// DeleteUpload deletes an upload.
func (h *Handler) DeleteUpload(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "upload deleted successfully",
	})
}

func toResponse(upload *Upload) UploadResponse {
	return UploadResponse{
		ID:              upload.ID,
		UserID:          upload.UserID,
		AppointmentID:   upload.AppointmentID,
		MedicalRecordID: upload.MedicalRecordID,
		HospitalID:      upload.HospitalID,
		FileName:        upload.FileName,
		OriginalName:    upload.OriginalName,
		FileType:        upload.FileType,
		FileSize:        upload.FileSize,
		FilePath:        upload.FilePath,
		Description:     upload.Description,
		IsPublic:        upload.IsPublic,
		CreatedAt:       upload.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       upload.UpdatedAt.Format(time.RFC3339),
	}
}
