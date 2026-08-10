package lab_requests

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	requests []LabRequest
	nextID   uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		requests: make([]LabRequest, 0),
		nextID:   1,
	}
}

// Create stores a lab request.
func (m *MockRepository) Create(request *LabRequest) error {
	request.ID = m.nextID
	m.nextID++

	m.requests = append(m.requests, *request)
	return nil
}

// GetByID retrieves a lab request by ID.
func (m *MockRepository) GetByID(id uint) (*LabRequest, error) {
	for _, request := range m.requests {
		if request.ID == id {
			r := request
			return &r, nil
		}
	}

	return nil, errors.New("lab request not found")
}

// GetAll retrieves all lab requests.
func (m *MockRepository) GetAll() ([]LabRequest, error) {
	return m.requests, nil
}

// GetByPatientID retrieves lab requests for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]LabRequest, error) {
	var requests []LabRequest

	for _, request := range m.requests {
		if request.PatientID == patientID {
			requests = append(requests, request)
		}
	}

	return requests, nil
}

// GetByDoctorID retrieves lab requests created by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]LabRequest, error) {
	var requests []LabRequest

	for _, request := range m.requests {
		if request.DoctorID == doctorID {
			requests = append(requests, request)
		}
	}

	return requests, nil
}

// GetByAppointmentID retrieves lab requests for an appointment.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]LabRequest, error) {
	var requests []LabRequest

	for _, request := range m.requests {
		if request.AppointmentID == appointmentID {
			requests = append(requests, request)
		}
	}

	return requests, nil
}

// GetByStatus retrieves lab requests by status.
func (m *MockRepository) GetByStatus(status string) ([]LabRequest, error) {
	var requests []LabRequest

	for _, request := range m.requests {
		if request.Status == status {
			requests = append(requests, request)
		}
	}

	return requests, nil
}

// Update updates a lab request.
func (m *MockRepository) Update(updated *LabRequest) error {
	for i, request := range m.requests {
		if request.ID == updated.ID {
			m.requests[i] = *updated
			return nil
		}
	}

	return errors.New("lab request not found")
}

// Delete removes a lab request.
func (m *MockRepository) Delete(id uint) error {
	for i, request := range m.requests {
		if request.ID == id {
			m.requests = append(m.requests[:i], m.requests[i+1:]...)
			return nil
		}
	}

	return errors.New("lab request not found")
}
