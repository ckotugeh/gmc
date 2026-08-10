package lab_results

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	results []LabResult
	nextID  uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		results: make([]LabResult, 0),
		nextID:  1,
	}
}

// Create stores a lab result.
func (m *MockRepository) Create(result *LabResult) error {
	result.ID = m.nextID
	m.nextID++

	m.results = append(m.results, *result)
	return nil
}

// GetByID retrieves a lab result by ID.
func (m *MockRepository) GetByID(id uint) (*LabResult, error) {
	for _, result := range m.results {
		if result.ID == id {
			r := result
			return &r, nil
		}
	}

	return nil, errors.New("lab result not found")
}

// GetAll retrieves all lab results.
func (m *MockRepository) GetAll() ([]LabResult, error) {
	return m.results, nil
}

// GetByLabRequestID retrieves results for a lab request.
func (m *MockRepository) GetByLabRequestID(labRequestID uint) ([]LabResult, error) {
	var results []LabResult

	for _, result := range m.results {
		if result.LabRequestID == labRequestID {
			results = append(results, result)
		}
	}

	return results, nil
}

// GetByPatientID retrieves lab results for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]LabResult, error) {
	var results []LabResult

	for _, result := range m.results {
		if result.PatientID == patientID {
			results = append(results, result)
		}
	}

	return results, nil
}

// GetByDoctorID retrieves lab results recorded by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]LabResult, error) {
	var results []LabResult

	for _, result := range m.results {
		if result.DoctorID == doctorID {
			results = append(results, result)
		}
	}

	return results, nil
}

// GetByStatus retrieves lab results by status.
func (m *MockRepository) GetByStatus(status string) ([]LabResult, error) {
	var results []LabResult

	for _, result := range m.results {
		if result.Status == status {
			results = append(results, result)
		}
	}

	return results, nil
}

// Update updates a lab result.
func (m *MockRepository) Update(updated *LabResult) error {
	for i, result := range m.results {
		if result.ID == updated.ID {
			m.results[i] = *updated
			return nil
		}
	}

	return errors.New("lab result not found")
}

// Delete removes a lab result.
func (m *MockRepository) Delete(id uint) error {
	for i, result := range m.results {
		if result.ID == id {
			m.results = append(m.results[:i], m.results[i+1:]...)
			return nil
		}
	}

	return errors.New("lab result not found")
}
