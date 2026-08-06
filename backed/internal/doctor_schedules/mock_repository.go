package doctor_schedules

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	schedules []DoctorSchedule
	nextID    uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		schedules: make([]DoctorSchedule, 0),
		nextID:    1,
	}
}

// Create stores a doctor schedule.
func (m *MockRepository) Create(schedule *DoctorSchedule) error {
	schedule.ID = m.nextID
	m.nextID++

	m.schedules = append(m.schedules, *schedule)
	return nil
}

// GetByID retrieves a schedule by ID.
func (m *MockRepository) GetByID(id uint) (*DoctorSchedule, error) {
	for _, schedule := range m.schedules {
		if schedule.ID == id {
			s := schedule
			return &s, nil
		}
	}

	return nil, errors.New("doctor schedule not found")
}

// GetAll retrieves all schedules.
func (m *MockRepository) GetAll() ([]DoctorSchedule, error) {
	return m.schedules, nil
}

// GetByDoctorID retrieves schedules for a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]DoctorSchedule, error) {
	var schedules []DoctorSchedule

	for _, schedule := range m.schedules {
		if schedule.DoctorID == doctorID {
			schedules = append(schedules, schedule)
		}
	}

	return schedules, nil
}

// GetByDay retrieves schedules for a given weekday.
func (m *MockRepository) GetByDay(day Weekday) ([]DoctorSchedule, error) {
	var schedules []DoctorSchedule

	for _, schedule := range m.schedules {
		if schedule.Day == day {
			schedules = append(schedules, schedule)
		}
	}

	return schedules, nil
}

// Update updates an existing schedule.
func (m *MockRepository) Update(updated *DoctorSchedule) error {
	for i, schedule := range m.schedules {
		if schedule.ID == updated.ID {
			m.schedules[i] = *updated
			return nil
		}
	}

	return errors.New("doctor schedule not found")
}

// Delete removes a schedule.
func (m *MockRepository) Delete(id uint) error {
	for i, schedule := range m.schedules {
		if schedule.ID == id {
			m.schedules = append(
				m.schedules[:i],
				m.schedules[i+1:]...,
			)
			return nil
		}
	}

	return errors.New("doctor schedule not found")
}
