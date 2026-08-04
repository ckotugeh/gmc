package video_consultations

// MockRepository is an in-memory implementation of Repository for testing.
type MockRepository struct {
	consultations []VideoConsultation
	nextID        uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		consultations: []VideoConsultation{},
		nextID:        1,
	}
}

// Create creates a new consultation.
func (m *MockRepository) Create(consultation *VideoConsultation) error {
	consultation.ID = m.nextID
	m.nextID++
	m.consultations = append(m.consultations, *consultation)
	return nil
}

// GetByID retrieves a consultation by ID.
func (m *MockRepository) GetByID(id uint) (*VideoConsultation, error) {
	for _, consultation := range m.consultations {
		if consultation.ID == id {
			c := consultation
			return &c, nil
		}
	}
	return nil, ErrConsultationNotFound
}

// GetByAppointmentID retrieves a consultation by appointment ID.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) (*VideoConsultation, error) {
	for _, consultation := range m.consultations {
		if consultation.AppointmentID == appointmentID {
			c := consultation
			return &c, nil
		}
	}
	return nil, ErrConsultationNotFound
}

// GetByDoctorID retrieves all consultations for a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]VideoConsultation, error) {
	var consultations []VideoConsultation

	for _, consultation := range m.consultations {
		if consultation.DoctorID == doctorID {
			consultations = append(consultations, consultation)
		}
	}

	return consultations, nil
}

// GetByPatientID retrieves all consultations for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]VideoConsultation, error) {
	var consultations []VideoConsultation

	for _, consultation := range m.consultations {
		if consultation.PatientID == patientID {
			consultations = append(consultations, consultation)
		}
	}

	return consultations, nil
}

// GetByRoomID retrieves a consultation by room ID.
func (m *MockRepository) GetByRoomID(roomID string) (*VideoConsultation, error) {
	for _, consultation := range m.consultations {
		if consultation.RoomID == roomID {
			c := consultation
			return &c, nil
		}
	}
	return nil, ErrConsultationNotFound
}

// GetAll retrieves all consultations.
func (m *MockRepository) GetAll() ([]VideoConsultation, error) {
	return m.consultations, nil
}

// Update updates a consultation.
func (m *MockRepository) Update(updated *VideoConsultation) error {
	for i, consultation := range m.consultations {
		if consultation.ID == updated.ID {
			m.consultations[i] = *updated
			return nil
		}
	}
	return ErrConsultationNotFound
}

// Delete deletes a consultation.
func (m *MockRepository) Delete(id uint) error {
	for i, consultation := range m.consultations {
		if consultation.ID == id {
			m.consultations = append(m.consultations[:i], m.consultations[i+1:]...)
			return nil
		}
	}
	return ErrConsultationNotFound
}
