package diagnoses

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	diagnoses []Diagnosis
	nextID    uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		diagnoses: make([]Diagnosis, 0),
		nextID:    1,
	}
}

// Create stores a diagnosis.
func (m *MockRepository) Create(diagnosis *Diagnosis) error {
	diagnosis.ID = m.nextID
	m.nextID++

	m.diagnoses = append(m.diagnoses, *diagnosis)
	return nil
}

// GetByID retrieves a diagnosis by ID.
func (m *MockRepository) GetByID(id uint) (*Diagnosis, error) {
	for _, diagnosis := range m.diagnoses {
		if diagnosis.ID == id {
			d := diagnosis
			return &d, nil
		}
	}

	return nil, errors.New("diagnosis not found")
}

// GetByAppointmentID retrieves diagnoses for an appointment.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	for _, diagnosis := range m.diagnoses {
		if diagnosis.AppointmentID == appointmentID {
			diagnoses = append(diagnoses, diagnosis)
		}
	}

	return diagnoses, nil
}

// GetByDoctorID retrieves diagnoses created by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	for _, diagnosis := range m.diagnoses {
		if diagnosis.DoctorID == doctorID {
			diagnoses = append(diagnoses, diagnosis)
		}
	}

	return diagnoses, nil
}

// GetByPatientID retrieves diagnoses for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	for _, diagnosis := range m.diagnoses {
		if diagnosis.PatientID == patientID {
			diagnoses = append(diagnoses, diagnosis)
		}
	}

	return diagnoses, nil
}

// GetByStatus retrieves diagnoses by status.
func (m *MockRepository) GetByStatus(status string) ([]Diagnosis, error) {
	var diagnoses []Diagnosis

	for _, diagnosis := range m.diagnoses {
		if diagnosis.Status == status {
			diagnoses = append(diagnoses, diagnosis)
		}
	}

	return diagnoses, nil
}

// GetAll retrieves all diagnoses.
func (m *MockRepository) GetAll() ([]Diagnosis, error) {
	return m.diagnoses, nil
}

// Update updates a diagnosis.
func (m *MockRepository) Update(updated *Diagnosis) error {
	for i, diagnosis := range m.diagnoses {
		if diagnosis.ID == updated.ID {
			m.diagnoses[i] = *updated
			return nil
		}
	}

	return errors.New("diagnosis not found")
}

// Delete removes a diagnosis.
func (m *MockRepository) Delete(id uint) error {
	for i, diagnosis := range m.diagnoses {
		if diagnosis.ID == id {
			m.diagnoses = append(m.diagnoses[:i], m.diagnoses[i+1:]...)
			return nil
		}
	}

	return errors.New("diagnosis not found")
}
