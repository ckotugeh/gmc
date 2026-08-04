package vitals

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	vitals []Vital
	nextID uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		vitals: make([]Vital, 0),
		nextID: 1,
	}
}

// Create stores a vital record.
func (m *MockRepository) Create(vital *Vital) error {
	vital.ID = m.nextID
	m.nextID++

	m.vitals = append(m.vitals, *vital)
	return nil
}

// GetByID retrieves a vital by ID.
func (m *MockRepository) GetByID(id uint) (*Vital, error) {
	for _, vital := range m.vitals {
		if vital.ID == id {
			v := vital
			return &v, nil
		}
	}

	return nil, errors.New("vital record not found")
}

// GetAll retrieves all vital records.
func (m *MockRepository) GetAll() ([]Vital, error) {
	return m.vitals, nil
}

// GetByPatientID retrieves all vitals for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]Vital, error) {
	var vitals []Vital

	for _, vital := range m.vitals {
		if vital.PatientID == patientID {
			vitals = append(vitals, vital)
		}
	}

	return vitals, nil
}

// GetByDoctorID retrieves all vitals recorded by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Vital, error) {
	var vitals []Vital

	for _, vital := range m.vitals {
		if vital.DoctorID == doctorID {
			vitals = append(vitals, vital)
		}
	}

	return vitals, nil
}

// GetByAppointmentID retrieves all vitals for an appointment.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]Vital, error) {
	var vitals []Vital

	for _, vital := range m.vitals {
		if vital.AppointmentID == appointmentID {
			vitals = append(vitals, vital)
		}
	}

	return vitals, nil
}

// Update updates a vital record.
func (m *MockRepository) Update(updated *Vital) error {
	for i, vital := range m.vitals {
		if vital.ID == updated.ID {
			m.vitals[i] = *updated
			return nil
		}
	}

	return errors.New("vital record not found")
}

// Delete removes a vital record.
func (m *MockRepository) Delete(id uint) error {
	for i, vital := range m.vitals {
		if vital.ID == id {
			m.vitals = append(m.vitals[:i], m.vitals[i+1:]...)
			return nil
		}
	}

	return errors.New("vital record not found")
}
