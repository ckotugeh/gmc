package clinical_notes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler handles clinical note HTTP requests.
type Handler struct {
	service Service
}

// NewHandler creates a new clinical note handler.
func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateClinicalNote creates a new clinical note.
func (h *Handler) CreateClinicalNote(c *gin.Context) {
	var req CreateClinicalNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := h.service.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, note)
}

// GetClinicalNote returns a clinical note by ID.
func (h *Handler) GetClinicalNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid clinical note id"})
		return
	}

	note, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

// GetAllClinicalNotes returns all clinical notes.
func (h *Handler) GetAllClinicalNotes(c *gin.Context) {
	notes, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetClinicalNotesByAppointment returns notes for an appointment.
func (h *Handler) GetClinicalNotesByAppointment(c *gin.Context) {
	appointmentID, err := strconv.ParseUint(c.Param("appointment_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
		return
	}

	notes, err := h.service.GetByAppointmentID(uint(appointmentID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetClinicalNotesByDoctor returns notes created by a doctor.
func (h *Handler) GetClinicalNotesByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor id"})
		return
	}

	notes, err := h.service.GetByDoctorID(uint(doctorID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetClinicalNotesByPatient returns notes for a patient.
func (h *Handler) GetClinicalNotesByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patient_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient id"})
		return
	}

	notes, err := h.service.GetByPatientID(uint(patientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetClinicalNotesByDiagnosis returns notes linked to a diagnosis.
func (h *Handler) GetClinicalNotesByDiagnosis(c *gin.Context) {
	diagnosisID, err := strconv.ParseUint(c.Param("diagnosis_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid diagnosis id"})
		return
	}

	notes, err := h.service.GetByDiagnosisID(uint(diagnosisID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// GetConfidentialClinicalNotes returns all confidential clinical notes.
func (h *Handler) GetConfidentialClinicalNotes(c *gin.Context) {
	notes, err := h.service.GetConfidential()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

// UpdateClinicalNote updates a clinical note.
func (h *Handler) UpdateClinicalNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid clinical note id"})
		return
	}

	var req UpdateClinicalNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note, err := h.service.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

// DeleteClinicalNote deletes a clinical note.
func (h *Handler) DeleteClinicalNote(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid clinical note id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "clinical note deleted successfully",
	})
}
