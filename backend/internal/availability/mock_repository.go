package availability

import (
	"errors"
	"time"
)

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	slots  []Availability
	nextID uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		slots:  make([]Availability, 0),
		nextID: 1,
	}
}

// Create stores a new availability slot.
func (m *MockRepository) Create(slot *Availability) error {
	slot.ID = m.nextID
	m.nextID++

	m.slots = append(m.slots, *slot)
	return nil
}

// GetByID retrieves an availability slot by ID.
func (m *MockRepository) GetByID(id uint) (*Availability, error) {
	for _, slot := range m.slots {
		if slot.ID == id {
			s := slot
			return &s, nil
		}
	}

	return nil, errors.New("availability slot not found")
}

// GetAll retrieves all availability slots.
func (m *MockRepository) GetAll() ([]Availability, error) {
	return m.slots, nil
}

// GetByDoctorID retrieves slots for a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Availability, error) {
	var slots []Availability

	for _, slot := range m.slots {
		if slot.DoctorID == doctorID {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

// GetByScheduleID retrieves slots for a schedule.
func (m *MockRepository) GetByScheduleID(scheduleID uint) ([]Availability, error) {
	var slots []Availability

	for _, slot := range m.slots {
		if slot.ScheduleID == scheduleID {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

// GetByDate retrieves slots for a specific date.
func (m *MockRepository) GetByDate(date time.Time) ([]Availability, error) {
	var slots []Availability

	for _, slot := range m.slots {
		if slot.Date.Equal(date) {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

// GetByDoctorAndDate retrieves slots for a doctor on a specific date.
func (m *MockRepository) GetByDoctorAndDate(doctorID uint, date time.Time) ([]Availability, error) {
	var slots []Availability

	for _, slot := range m.slots {
		if slot.DoctorID == doctorID && slot.Date.Equal(date) {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

// Update updates an availability slot.
func (m *MockRepository) Update(updated *Availability) error {
	for i, slot := range m.slots {
		if slot.ID == updated.ID {
			m.slots[i] = *updated
			return nil
		}
	}

	return errors.New("availability slot not found")
}

// Delete removes an availability slot.
func (m *MockRepository) Delete(id uint) error {
	for i, slot := range m.slots {
		if slot.ID == id {
			m.slots = append(m.slots[:i], m.slots[i+1:]...)
			return nil
		}
	}

	return errors.New("availability slot not found")
}
