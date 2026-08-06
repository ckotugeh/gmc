package doctor_reviews

import "errors"

// MockRepository is an in-memory implementation of Repository.
type MockRepository struct {
	reviews []DoctorReview
	nextID  uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		reviews: make([]DoctorReview, 0),
		nextID:  1,
	}
}

// Create stores a doctor review.
func (m *MockRepository) Create(review *DoctorReview) error {
	for _, r := range m.reviews {
		if r.AppointmentID == review.AppointmentID {
			return errors.New("review already exists for this appointment")
		}
	}

	review.ID = m.nextID
	m.nextID++

	m.reviews = append(m.reviews, *review)
	return nil
}

// GetByID retrieves a review by ID.
func (m *MockRepository) GetByID(id uint) (*DoctorReview, error) {
	for _, review := range m.reviews {
		if review.ID == id {
			r := review
			return &r, nil
		}
	}

	return nil, errors.New("doctor review not found")
}

// GetByAppointmentID retrieves a review by appointment ID.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) (*DoctorReview, error) {
	for _, review := range m.reviews {
		if review.AppointmentID == appointmentID {
			r := review
			return &r, nil
		}
	}

	return nil, errors.New("doctor review not found")
}

// GetByDoctorID retrieves all reviews for a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]DoctorReview, error) {
	var reviews []DoctorReview

	for _, review := range m.reviews {
		if review.DoctorID == doctorID {
			reviews = append(reviews, review)
		}
	}

	return reviews, nil
}

// GetByPatientID retrieves all reviews created by a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]DoctorReview, error) {
	var reviews []DoctorReview

	for _, review := range m.reviews {
		if review.PatientID == patientID {
			reviews = append(reviews, review)
		}
	}

	return reviews, nil
}

// GetPublishedByDoctorID retrieves all published reviews for a doctor.
func (m *MockRepository) GetPublishedByDoctorID(doctorID uint) ([]DoctorReview, error) {
	var reviews []DoctorReview

	for _, review := range m.reviews {
		if review.DoctorID == doctorID && review.IsPublished {
			reviews = append(reviews, review)
		}
	}

	return reviews, nil
}

// GetAll retrieves all doctor reviews.
func (m *MockRepository) GetAll() ([]DoctorReview, error) {
	return m.reviews, nil
}

// Update updates a doctor review.
func (m *MockRepository) Update(updated *DoctorReview) error {
	for i, review := range m.reviews {
		if review.ID == updated.ID {
			m.reviews[i] = *updated
			return nil
		}
	}

	return errors.New("doctor review not found")
}

// Delete removes a doctor review.
func (m *MockRepository) Delete(id uint) error {
	for i, review := range m.reviews {
		if review.ID == id {
			m.reviews = append(m.reviews[:i], m.reviews[i+1:]...)
			return nil
		}
	}

	return errors.New("doctor review not found")
}
