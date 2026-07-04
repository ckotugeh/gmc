package medicalrecords

// MockRepository is an in-memory implementation of Repository for testing.
type MockRepository struct {
	records []MedicalRecord
	nextID  uint
}

// NewMockRepository creates a new mock repository.
func NewMockRepository() Repository {
	return &MockRepository{
		records: []MedicalRecord{},
		nextID:  1,
	}
}

// Create creates a new medical record.
func (m *MockRepository) Create(record *MedicalRecord) error {
	record.ID = m.nextID
	m.nextID++
	m.records = append(m.records, *record)
	return nil
}

// GetByID retrieves a medical record by ID.
func (m *MockRepository) GetByID(id uint) (*MedicalRecord, error) {
	for _, record := range m.records {
		if record.ID == id {
			r := record
			return &r, nil
		}
	}
	return nil, ErrMedicalRecordNotFound
}

// GetAll retrieves all medical records.
func (m *MockRepository) GetAll() ([]MedicalRecord, error) {
	return m.records, nil
}

// GetByPatientID retrieves all medical records for a patient.
func (m *MockRepository) GetByPatientID(patientID uint) ([]MedicalRecord, error) {
	var records []MedicalRecord
	for _, record := range m.records {
		if record.PatientID == patientID {
			records = append(records, record)
		}
	}
	return records, nil
}

// GetByDoctorID retrieves all medical records created by a doctor.
func (m *MockRepository) GetByDoctorID(doctorID uint) ([]MedicalRecord, error) {
	var records []MedicalRecord
	for _, record := range m.records {
		if record.DoctorID == doctorID {
			records = append(records, record)
		}
	}
	return records, nil
}

// Update updates a medical record.
func (m *MockRepository) Update(updated *MedicalRecord) error {
	for i, record := range m.records {
		if record.ID == updated.ID {
			m.records[i] = *updated
			return nil
		}
	}
	return ErrMedicalRecordNotFound
}

// Delete deletes a medical record.
func (m *MockRepository) Delete(id uint) error {
	for i, record := range m.records {
		if record.ID == id {
			m.records = append(m.records[:i], m.records[i+1:]...)
			return nil
		}
	}
	return ErrMedicalRecordNotFound
}
