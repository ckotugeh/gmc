package prescriptions

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	prescriptions []Prescription
	nextID        uint
	nextItemID    uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		prescriptions: make([]Prescription, 0),
		nextID:        1,
		nextItemID:    1,
	}
}

// Create stores a prescription.
func (m *MockRepository) Create(prescription *Prescription) error {
	prescription.ID = m.nextID
	m.nextID++

	for i := range prescription.Items {
		prescription.Items[i].ID = m.nextItemID
		prescription.Items[i].PrescriptionID = prescription.ID
		m.nextItemID++
	}

	m.prescriptions = append(m.prescriptions, *prescription)
	return nil
}

// GetByID retrieves a prescription by ID.
func (m *MockRepository) GetByID(id uint) (*Prescription, error) {
	for _, prescription := range m.prescriptions {
		if prescription.ID == id {
			p := prescription
			return &p, nil
		}
	}

	return nil, errors.New("prescription not found")
}

// GetAll retrieves all prescriptions.
func (m *MockRepository) GetAll() ([]Prescription, error) {
	return m.prescriptions, nil
}

// GetByDoctorID retrieves prescriptions by doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Prescription, error) {
	var prescriptions []Prescription

	for _, prescription := range m.prescriptions {
		if prescription.DoctorID == doctorID {
			prescriptions = append(prescriptions, prescription)
		}
	}

	return prescriptions, nil
}

// GetByPatientID retrieves prescriptions by patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]Prescription, error) {
	var prescriptions []Prescription

	for _, prescription := range m.prescriptions {
		if prescription.PatientID == patientID {
			prescriptions = append(prescriptions, prescription)
		}
	}

	return prescriptions, nil
}

// GetByAppointmentID retrieves prescriptions by appointment.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]Prescription, error) {
	var prescriptions []Prescription

	for _, prescription := range m.prescriptions {
		if prescription.AppointmentID == appointmentID {
			prescriptions = append(prescriptions, prescription)
		}
	}

	return prescriptions, nil
}

// Update updates an existing prescription.
func (m *MockRepository) Update(updated *Prescription) error {
	for i, prescription := range m.prescriptions {
		if prescription.ID == updated.ID {
			for j := range updated.Items {
				if updated.Items[j].ID == 0 {
					updated.Items[j].ID = m.nextItemID
					updated.Items[j].PrescriptionID = updated.ID
					m.nextItemID++
				}
			}

			m.prescriptions[i] = *updated
			return nil
		}
	}

	return errors.New("prescription not found")
}

// Delete removes a prescription.
func (m *MockRepository) Delete(id uint) error {
	for i, prescription := range m.prescriptions {
		if prescription.ID == id {
			m.prescriptions = append(
				m.prescriptions[:i],
				m.prescriptions[i+1:]...,
			)
			return nil
		}
	}

	return errors.New("prescription not found")
}
