package allergies

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	allergies []Allergy
	nextID    uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		allergies: make([]Allergy, 0),
		nextID:    1,
	}
}

// Create stores an allergy record.
func (m *MockRepository) Create(allergy *Allergy) error {
	allergy.ID = m.nextID
	m.nextID++

	m.allergies = append(m.allergies, *allergy)
	return nil
}

// GetByID retrieves an allergy by ID.
func (m *MockRepository) GetByID(id uint) (*Allergy, error) {
	for _, allergy := range m.allergies {
		if allergy.ID == id {
			a := allergy
			return &a, nil
		}
	}

	return nil, errors.New("allergy record not found")
}

// GetAll retrieves all allergy records.
func (m *MockRepository) GetAll() ([]Allergy, error) {
	return m.allergies, nil
}

// GetByPatientID retrieves allergies for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]Allergy, error) {
	var allergies []Allergy

	for _, allergy := range m.allergies {
		if allergy.PatientID == patientID {
			allergies = append(allergies, allergy)
		}
	}

	return allergies, nil
}

// GetByDoctorID retrieves allergies recorded by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Allergy, error) {
	var allergies []Allergy

	for _, allergy := range m.allergies {
		if allergy.DoctorID == doctorID {
			allergies = append(allergies, allergy)
		}
	}

	return allergies, nil
}

// GetBySeverity retrieves allergies by severity.
func (m *MockRepository) GetBySeverity(severity string) ([]Allergy, error) {
	var allergies []Allergy

	for _, allergy := range m.allergies {
		if allergy.Severity == severity {
			allergies = append(allergies, allergy)
		}
	}

	return allergies, nil
}

// GetActive retrieves active allergies.
func (m *MockRepository) GetActive() ([]Allergy, error) {
	var allergies []Allergy

	for _, allergy := range m.allergies {
		if allergy.Status == "Active" {
			allergies = append(allergies, allergy)
		}
	}

	return allergies, nil
}

// Update updates an allergy record.
func (m *MockRepository) Update(updated *Allergy) error {
	for i, allergy := range m.allergies {
		if allergy.ID == updated.ID {
			m.allergies[i] = *updated
			return nil
		}
	}

	return errors.New("allergy record not found")
}

// Delete removes an allergy record.
func (m *MockRepository) Delete(id uint) error {
	for i, allergy := range m.allergies {
		if allergy.ID == id {
			m.allergies = append(m.allergies[:i], m.allergies[i+1:]...)
			return nil
		}
	}

	return errors.New("allergy record not found")
}
