package medical_specialties

import (
	"errors"
	"strings"
)

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	specialties []MedicalSpecialty
	nextID      uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		specialties: make([]MedicalSpecialty, 0),
		nextID:      1,
	}
}

// Create stores a medical specialty.
func (m *MockRepository) Create(specialty *MedicalSpecialty) error {
	for _, s := range m.specialties {
		if strings.EqualFold(s.Name, specialty.Name) {
			return errors.New("medical specialty name already exists")
		}
		if strings.EqualFold(s.Code, specialty.Code) {
			return errors.New("medical specialty code already exists")
		}
	}

	specialty.ID = m.nextID
	m.nextID++

	m.specialties = append(m.specialties, *specialty)
	return nil
}

// GetByID retrieves a specialty by ID.
func (m *MockRepository) GetByID(id uint) (*MedicalSpecialty, error) {
	for _, specialty := range m.specialties {
		if specialty.ID == id {
			s := specialty
			return &s, nil
		}
	}

	return nil, errors.New("medical specialty not found")
}

// GetByName retrieves a specialty by name.
func (m *MockRepository) GetByName(name string) (*MedicalSpecialty, error) {
	for _, specialty := range m.specialties {
		if strings.EqualFold(specialty.Name, name) {
			s := specialty
			return &s, nil
		}
	}

	return nil, errors.New("medical specialty not found")
}

// GetByCode retrieves a specialty by code.
func (m *MockRepository) GetByCode(code string) (*MedicalSpecialty, error) {
	for _, specialty := range m.specialties {
		if strings.EqualFold(specialty.Code, code) {
			s := specialty
			return &s, nil
		}
	}

	return nil, errors.New("medical specialty not found")
}

// GetAll retrieves all specialties.
func (m *MockRepository) GetAll() ([]MedicalSpecialty, error) {
	return m.specialties, nil
}

// GetActive retrieves all active specialties.
func (m *MockRepository) GetActive() ([]MedicalSpecialty, error) {
	var active []MedicalSpecialty

	for _, specialty := range m.specialties {
		if specialty.IsActive {
			active = append(active, specialty)
		}
	}

	return active, nil
}

// Update updates a specialty.
func (m *MockRepository) Update(updated *MedicalSpecialty) error {
	for i, specialty := range m.specialties {
		if specialty.ID == updated.ID {
			m.specialties[i] = *updated
			return nil
		}
	}

	return errors.New("medical specialty not found")
}

// Delete removes a specialty.
func (m *MockRepository) Delete(id uint) error {
	for i, specialty := range m.specialties {
		if specialty.ID == id {
			m.specialties = append(
				m.specialties[:i],
				m.specialties[i+1:]...,
			)
			return nil
		}
	}

	return errors.New("medical specialty not found")
}
