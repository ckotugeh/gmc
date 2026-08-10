package hospitals

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Handler handles hospital HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new hospital handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateHospital creates a new hospital.
func (h *Handler) CreateHospital(c *gin.Context) {
	var req CreateHospitalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital, err := h.service.Create(req)
	if err != nil {
		switch err {
		case ErrInvalidHospital,
			ErrHospitalEmailExists,
			ErrHospitalLicenseExists:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create hospital"})
		}
		return
	}

	c.JSON(http.StatusCreated, toHospitalResponse(hospital))
}

// GetHospital returns a hospital by ID.
func (h *Handler) GetHospital(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hospital id"})
		return
	}

	hospital, err := h.service.GetByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || err == ErrHospitalNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get hospital"})
		return
	}

	c.JSON(http.StatusOK, toHospitalResponse(hospital))
}

// GetHospitals returns all hospitals.
func (h *Handler) GetHospitals(c *gin.Context) {
	hospitals, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve hospitals"})
		return
	}

	response := make([]HospitalResponse, 0, len(hospitals))
	for _, hospital := range hospitals {
		response = append(response, toHospitalResponse(&hospital))
	}

	c.JSON(http.StatusOK, response)
}

// UpdateHospital updates a hospital.
func (h *Handler) UpdateHospital(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hospital id"})
		return
	}

	var req UpdateHospitalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital, err := h.service.Update(uint(id), req)
	if err != nil {
		switch err {
		case ErrHospitalNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case ErrHospitalEmailExists,
			ErrHospitalLicenseExists:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update hospital"})
		}
		return
	}

	c.JSON(http.StatusOK, toHospitalResponse(hospital))
}

// DeleteHospital deletes a hospital.
func (h *Handler) DeleteHospital(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid hospital id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		if err == ErrHospitalNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete hospital"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "hospital deleted successfully",
	})
}

func toHospitalResponse(h *Hospital) HospitalResponse {
	return HospitalResponse{
		ID:            h.ID,
		Name:          h.Name,
		Description:   h.Description,
		Email:         h.Email,
		Phone:         h.Phone,
		Website:       h.Website,
		Address:       h.Address,
		City:          h.City,
		State:         h.State,
		Country:       h.Country,
		ZipCode:       h.ZipCode,
		LicenseNumber: h.LicenseNumber,
		IsActive:      h.IsActive,
		CreatedAt:     h.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     h.UpdatedAt.Format(time.RFC3339),
	}
}
