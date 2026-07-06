package diagnoses

import "testing"

func TestCreateDiagnosis(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		DiagnosisCode: "J11.1",
		Condition:     "Influenza",
		Description:   "Seasonal influenza",
		Severity:      "Moderate",
		Status:        "Active",
		Notes:         "Patient advised to rest.",
	}

	diagnosis, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if diagnosis.ID == 0 {
		t.Fatal("expected diagnosis ID to be assigned")
	}

	if diagnosis.Condition != req.Condition {
		t.Fatalf("expected %s, got %s", req.Condition, diagnosis.Condition)
	}
}

func TestCreateDiagnosisValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateDiagnosisRequest{
		{},
		{
			DoctorID:  1,
			PatientID: 1,
			Condition: "Influenza",
		},
		{
			AppointmentID: 1,
			PatientID:     1,
			Condition:     "Influenza",
		},
		{
			AppointmentID: 1,
			DoctorID:      1,
			Condition:     "Influenza",
		},
		{
			AppointmentID: 1,
			DoctorID:      1,
			PatientID:     1,
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidDiagnosis {
			t.Fatalf("expected %v, got %v", ErrInvalidDiagnosis, err)
		}
	}
}

func TestGetDiagnosis(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Condition:     "Influenza",
	})

	diagnosis, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if diagnosis.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, diagnosis.ID)
	}
}

func TestGetByAppointmentID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDiagnosisRequest{
		AppointmentID: 10,
		DoctorID:      1,
		PatientID:     1,
		Condition:     "Malaria",
	})

	diagnoses, err := service.GetByAppointmentID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(diagnoses) != 1 {
		t.Fatalf("expected 1 diagnosis, got %d", len(diagnoses))
	}
}

func TestGetByDoctorID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     1,
		Condition:     "Asthma",
	})

	diagnoses, err := service.GetByDoctorID(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(diagnoses) != 1 {
		t.Fatalf("expected 1 diagnosis, got %d", len(diagnoses))
	}
}

func TestGetByPatientID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     5,
		Condition:     "Diabetes",
	})

	diagnoses, err := service.GetByPatientID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(diagnoses) != 1 {
		t.Fatalf("expected 1 diagnosis, got %d", len(diagnoses))
	}
}

func TestGetByStatus(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Condition:     "Hypertension",
		Status:        "Resolved",
	})

	diagnoses, err := service.GetByStatus("Resolved")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(diagnoses) != 1 {
		t.Fatalf("expected 1 diagnosis, got %d", len(diagnoses))
	}
}

func TestUpdateDiagnosis(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Condition:     "Influenza",
		Status:        "Active",
	})

	updated, err := service.Update(created.ID, UpdateDiagnosisRequest{
		Condition: "COVID-19",
		Status:    "Resolved",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Condition != "COVID-19" {
		t.Fatal("condition was not updated")
	}

	if updated.Status != "Resolved" {
		t.Fatal("status was not updated")
	}
}

func TestDeleteDiagnosis(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDiagnosisRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Condition:     "Influenza",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrDiagnosisNotFound {
		t.Fatalf("expected %v, got %v", ErrDiagnosisNotFound, err)
	}
}
