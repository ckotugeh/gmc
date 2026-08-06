package medicalrecords

import "testing"

func TestCreateMedicalRecord(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateMedicalRecordRequest{
		PatientID:    1,
		DoctorID:     2,
		Diagnosis:    "Malaria",
		Symptoms:     "Fever, headache",
		Treatment:    "ACT",
		Prescription: "Artemether-Lumefantrine",
		Notes:        "Patient to return in 7 days",
	}

	record, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if record.ID == 0 {
		t.Error("expected ID to be assigned")
	}

	if record.Diagnosis != req.Diagnosis {
		t.Errorf("expected diagnosis %q, got %q", req.Diagnosis, record.Diagnosis)
	}
}

func TestCreateMedicalRecordValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  2,
	}

	_, err := service.Create(req)
	if err == nil {
		t.Fatal("expected validation error")
	}

	if err != ErrInvalidMedicalRecord {
		t.Fatalf("expected ErrInvalidMedicalRecord, got %v", err)
	}
}

func TestGetMedicalRecord(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	record, _ := service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  2,
		Diagnosis: "Typhoid",
	})

	found, err := service.GetByID(record.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if found.ID != record.ID {
		t.Errorf("expected ID %d, got %d", record.ID, found.ID)
	}
}

func TestGetAllMedicalRecords(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, err := service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  2,
		Diagnosis: "Asthma",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.Create(CreateMedicalRecordRequest{
		PatientID: 2,
		DoctorID:  3,
		Diagnosis: "Diabetes",
	})

	records, err := service.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(records))
	}
}

func TestGetPatientMedicalRecords(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, err := service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  2,
		Diagnosis: "Malaria",
	})

	service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  3,
		Diagnosis: "Flu",
	})

	service.Create(CreateMedicalRecordRequest{
		PatientID: 2,
		DoctorID:  2,
		Diagnosis: "Cold",
	})

	records, err := service.GetByPatientID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(records))
	}
}

func TestGetDoctorMedicalRecords(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  5,
		Diagnosis: "Hypertension",
	})

	service.Create(CreateMedicalRecordRequest{
		PatientID: 2,
		DoctorID:  5,
		Diagnosis: "Diabetes",
	})

	service.Create(CreateMedicalRecordRequest{
		PatientID: 3,
		DoctorID:  8,
		Diagnosis: "Asthma",
	})

	records, err := service.GetByDoctorID(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(records) != 2 {
		t.Fatalf("expected 2 records, got %d", len(records))
	}
}

func TestUpdateMedicalRecord(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	record, _ := service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  2,
		Diagnosis: "Flu",
	})

	updated, err := service.Update(record.ID, UpdateMedicalRecordRequest{
		Diagnosis: "Severe Flu",
		Treatment: "Medication and rest",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updated.Diagnosis != "Severe Flu" {
		t.Errorf("expected diagnosis to be updated")
	}

	if updated.Treatment != "Medication and rest" {
		t.Errorf("expected treatment to be updated")
	}
}

func TestDeleteMedicalRecord(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	record, _ := service.Create(CreateMedicalRecordRequest{
		PatientID: 1,
		DoctorID:  2,
		Diagnosis: "Migraine",
	})

	err := service.Delete(record.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.GetByID(record.ID)
	if err == nil {
		t.Fatal("expected record to be deleted")
	}
}
