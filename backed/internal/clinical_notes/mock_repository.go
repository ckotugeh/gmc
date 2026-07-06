package clinical_notes

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	notes  []ClinicalNote
	nextID uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		notes:  make([]ClinicalNote, 0),
		nextID: 1,
	}
}

// Create stores a clinical note.
func (m *MockRepository) Create(note *ClinicalNote) error {
	note.ID = m.nextID
	m.nextID++

	m.notes = append(m.notes, *note)
	return nil
}

// GetByID retrieves a clinical note by ID.
func (m *MockRepository) GetByID(id uint) (*ClinicalNote, error) {
	for _, note := range m.notes {
		if note.ID == id {
			n := note
			return &n, nil
		}
	}

	return nil, errors.New("clinical note not found")
}

// GetByAppointmentID retrieves notes for an appointment.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	for _, note := range m.notes {
		if note.AppointmentID == appointmentID {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// GetByDoctorID retrieves notes created by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	for _, note := range m.notes {
		if note.DoctorID == doctorID {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// GetByPatientID retrieves notes for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	for _, note := range m.notes {
		if note.PatientID == patientID {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// GetByDiagnosisID retrieves notes for a diagnosis.
func (m *MockRepository) GetByDiagnosisID(diagnosisID uint) ([]ClinicalNote, error) {
	var notes []ClinicalNote

	for _, note := range m.notes {
		if note.DiagnosisID == diagnosisID {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// GetConfidential retrieves all confidential notes.
func (m *MockRepository) GetConfidential() ([]ClinicalNote, error) {
	var notes []ClinicalNote

	for _, note := range m.notes {
		if note.IsConfidential {
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// GetAll retrieves all clinical notes.
func (m *MockRepository) GetAll() ([]ClinicalNote, error) {
	return m.notes, nil
}

// Update updates a clinical note.
func (m *MockRepository) Update(updated *ClinicalNote) error {
	for i, note := range m.notes {
		if note.ID == updated.ID {
			m.notes[i] = *updated
			return nil
		}
	}

	return errors.New("clinical note not found")
}

// Delete removes a clinical note.
func (m *MockRepository) Delete(id uint) error {
	for i, note := range m.notes {
		if note.ID == id {
			m.notes = append(m.notes[:i], m.notes[i+1:]...)
			return nil
		}
	}

	return errors.New("clinical note not found")
}
