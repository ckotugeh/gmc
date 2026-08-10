package appointments

import (
	"errors"
	"sync"
)

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	mu           sync.RWMutex
	appointments map[uint]*Appointment
	nextID       uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() *MockRepository {
	return &MockRepository{
		appointments: make(map[uint]*Appointment),
		nextID:       1,
	}
}

// Create stores a new appointment.
func (m *MockRepository) Create(appointment *Appointment) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	appointment.ID = m.nextID
	m.nextID++

	copy := *appointment
	m.appointments[appointment.ID] = &copy

	return nil
}

// GetByID returns an appointment by ID.
func (m *MockRepository) GetByID(id uint) (*Appointment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	appointment, exists := m.appointments[id]
	if !exists {
		return nil, errors.New("appointment not found")
	}

	copy := *appointment
	return &copy, nil
}

// GetAll returns all appointments.
func (m *MockRepository) GetAll() ([]Appointment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	appointments := make([]Appointment, 0, len(m.appointments))

	for _, appointment := range m.appointments {
		appointments = append(appointments, *appointment)
	}

	return appointments, nil
}

// GetByDoctorID returns appointments for a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Appointment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var appointments []Appointment

	for _, appointment := range m.appointments {
		if appointment.DoctorID == doctorID {
			appointments = append(appointments, *appointment)
		}
	}

	return appointments, nil
}

// GetByPatientID returns appointments for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]Appointment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var appointments []Appointment

	for _, appointment := range m.appointments {
		if appointment.PatientID == patientID {
			appointments = append(appointments, *appointment)
		}
	}

	return appointments, nil
}

// Update updates an existing appointment.
func (m *MockRepository) Update(appointment *Appointment) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.appointments[appointment.ID]; !exists {
		return errors.New("appointment not found")
	}

	copy := *appointment
	m.appointments[appointment.ID] = &copy

	return nil
}

// Delete removes an appointment.
func (m *MockRepository) Delete(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.appointments[id]; !exists {
		return errors.New("appointment not found")
	}

	delete(m.appointments, id)
	return nil
}
