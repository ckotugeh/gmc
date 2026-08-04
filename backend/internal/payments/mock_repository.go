package payments

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	payments []Payment
	nextID   uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		payments: make([]Payment, 0),
		nextID:   1,
	}
}

// Create stores a payment.
func (m *MockRepository) Create(payment *Payment) error {
	payment.ID = m.nextID
	m.nextID++

	m.payments = append(m.payments, *payment)
	return nil
}

// GetByID retrieves a payment by ID.
func (m *MockRepository) GetByID(id uint) (*Payment, error) {
	for _, payment := range m.payments {
		if payment.ID == id {
			p := payment
			return &p, nil
		}
	}

	return nil, errors.New("payment not found")
}

// GetAll retrieves all payments.
func (m *MockRepository) GetAll() ([]Payment, error) {
	return m.payments, nil
}

// GetByAppointmentID retrieves payments by appointment.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]Payment, error) {
	var payments []Payment

	for _, payment := range m.payments {
		if payment.AppointmentID == appointmentID {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

// GetByPatientID retrieves payments by patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]Payment, error) {
	var payments []Payment

	for _, payment := range m.payments {
		if payment.PatientID == patientID {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

// GetByDoctorID retrieves payments by doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]Payment, error) {
	var payments []Payment

	for _, payment := range m.payments {
		if payment.DoctorID == doctorID {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

// GetByHospitalID retrieves payments by hospital.
func (m *MockRepository) GetByHospitalID(hospitalID uint) ([]Payment, error) {
	var payments []Payment

	for _, payment := range m.payments {
		if payment.HospitalID != nil && *payment.HospitalID == hospitalID {
			payments = append(payments, payment)
		}
	}

	return payments, nil
}

// Update updates an existing payment.
func (m *MockRepository) Update(updated *Payment) error {
	for i, payment := range m.payments {
		if payment.ID == updated.ID {
			m.payments[i] = *updated
			return nil
		}
	}

	return errors.New("payment not found")
}

// Delete removes a payment.
func (m *MockRepository) Delete(id uint) error {
	for i, payment := range m.payments {
		if payment.ID == id {
			m.payments = append(m.payments[:i], m.payments[i+1:]...)
			return nil
		}
	}

	return errors.New("payment not found")
}

// GetSummary returns payment statistics.
func (m *MockRepository) GetSummary() (*PaymentSummaryResponse, error) {
	summary := &PaymentSummaryResponse{}

	summary.TotalPayments = int64(len(m.payments))

	for _, payment := range m.payments {
		switch payment.Status {
		case StatusPending:
			summary.PendingAmount += payment.Amount
		case StatusPaid:
			summary.PaidAmount += payment.Amount
			summary.TotalRevenue += payment.Amount
		case StatusRefunded:
			summary.RefundedAmount += payment.Amount
		}
	}

	return summary, nil
}
