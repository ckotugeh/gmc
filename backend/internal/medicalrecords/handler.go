package medicalrecords

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

// NewHandler creates a new medical records handler.
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateMedicalRecord handles POST /medical-records.
func (h *Handler) CreateMedicalRecord(c *gin.Context) {
	var req CreateMedicalRecordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := h.service.Create(req)
	if err != nil {
		if errors.Is(err, ErrInvalidMedicalRecord) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

// GetMedicalRecord handles GET /medical-records/:id.
func (h *Handler) GetMedicalRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid medical record id"})
		return
	}

	record, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "medical record not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// GetMedicalRecords handles GET /medical-records.
func (h *Handler) GetMedicalRecords(c *gin.Context) {
	records, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// GetPatientMedicalRecords handles GET /medical-records/patient/:patientID.
func (h *Handler) GetPatientMedicalRecords(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patientID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	records, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// GetDoctorMedicalRecords handles GET /medical-records/doctor/:doctorID.
func (h *Handler) GetDoctorMedicalRecords(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctorID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	records, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, records)
}

// UpdateMedicalRecord handles PUT /medical-records/:id.
func (h *Handler) UpdateMedicalRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid medical record id"})
		return
	}

	var req UpdateMedicalRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	record, err := h.service.Update(uint(id), req)
	if err != nil {
		if errors.Is(err, ErrMedicalRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// DeleteMedicalRecord handles DELETE /medical-records/:id.
func (h *Handler) DeleteMedicalRecord(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid medical record id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		if errors.Is(err, ErrMedicalRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "medical record deleted successfully",
	})
}
