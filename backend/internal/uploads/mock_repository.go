package uploads

// MockRepository is an in-memory implementation of Repository for testing.
type MockRepository struct {
	uploads []Upload
	nextID  uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		uploads: []Upload{},
		nextID:  1,
	}
}

// Create creates a new upload.
func (m *MockRepository) Create(upload *Upload) error {
	upload.ID = m.nextID
	m.nextID++
	m.uploads = append(m.uploads, *upload)
	return nil
}

// GetByID retrieves an upload by ID.
func (m *MockRepository) GetByID(id uint) (*Upload, error) {
	for _, upload := range m.uploads {
		if upload.ID == id {
			u := upload
			return &u, nil
		}
	}
	return nil, ErrUploadNotFound
}

// GetByUserID retrieves uploads by user ID.
func (m *MockRepository) GetByUserID(userID uint) ([]Upload, error) {
	var uploads []Upload
	for _, upload := range m.uploads {
		if upload.UserID == userID {
			uploads = append(uploads, upload)
		}
	}
	return uploads, nil
}

// GetByAppointmentID retrieves uploads by appointment ID.
func (m *MockRepository) GetByAppointmentID(appointmentID uint) ([]Upload, error) {
	var uploads []Upload
	for _, upload := range m.uploads {
		if upload.AppointmentID != nil && *upload.AppointmentID == appointmentID {
			uploads = append(uploads, upload)
		}
	}
	return uploads, nil
}

// GetByMedicalRecordID retrieves uploads by medical record ID.
func (m *MockRepository) GetByMedicalRecordID(recordID uint) ([]Upload, error) {
	var uploads []Upload
	for _, upload := range m.uploads {
		if upload.MedicalRecordID != nil && *upload.MedicalRecordID == recordID {
			uploads = append(uploads, upload)
		}
	}
	return uploads, nil
}

// GetByHospitalID retrieves uploads by hospital ID.
func (m *MockRepository) GetByHospitalID(hospitalID uint) ([]Upload, error) {
	var uploads []Upload
	for _, upload := range m.uploads {
		if upload.HospitalID != nil && *upload.HospitalID == hospitalID {
			uploads = append(uploads, upload)
		}
	}
	return uploads, nil
}

// Update updates an upload.
func (m *MockRepository) Update(updated *Upload) error {
	for i, upload := range m.uploads {
		if upload.ID == updated.ID {
			m.uploads[i] = *updated
			return nil
		}
	}
	return ErrUploadNotFound
}

// Delete deletes an upload.
func (m *MockRepository) Delete(id uint) error {
	for i, upload := range m.uploads {
		if upload.ID == id {
			m.uploads = append(m.uploads[:i], m.uploads[i+1:]...)
			return nil
		}
	}
	return ErrUploadNotFound
}
