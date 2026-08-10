package vitals

import "testing"

func TestCreateVital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateVitalRequest{
		PatientID:        1,
		DoctorID:         1,
		AppointmentID:    1,
		Temperature:      36.8,
		HeartRate:        72,
		RespiratoryRate:  18,
		SystolicBP:       120,
		DiastolicBP:      80,
		OxygenSaturation: 98,
		Weight:           70.5,
		Height:           175,
		BMI:              23.0,
		Notes:            "Patient vitals normal",
	}

	vital, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if vital.ID == 0 {
		t.Fatal("expected vital ID to be assigned")
	}

	if vital.PatientID != req.PatientID {
		t.Fatalf("expected patient ID %d, got %d", req.PatientID, vital.PatientID)
	}
}

func TestCreateVitalValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateVitalRequest{
		{},
		{DoctorID: 1},
		{PatientID: 1},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidVital {
			t.Fatalf("expected %v, got %v", ErrInvalidVital, err)
		}
	}
}

func TestGetVitalByID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVitalRequest{
		PatientID: 1,
		DoctorID:  1,
	})

	vital, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if vital.ID != created.ID {
		t.Fatal("returned wrong vital")
	}
}

func TestGetAllVitals(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateVitalRequest{
		PatientID: 1,
		DoctorID:  1,
	})

	_, _ = service.Create(CreateVitalRequest{
		PatientID: 2,
		DoctorID:  1,
	})

	vitals, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(vitals) != 2 {
		t.Fatalf("expected 2 vitals, got %d", len(vitals))
	}
}

func TestGetVitalsByPatient(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateVitalRequest{
		PatientID: 5,
		DoctorID:  1,
	})

	vitals, err := service.GetByPatientID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(vitals) != 1 {
		t.Fatalf("expected 1 vital, got %d", len(vitals))
	}
}

func TestGetVitalsByDoctor(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateVitalRequest{
		PatientID: 1,
		DoctorID:  8,
	})

	vitals, err := service.GetByDoctorID(8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(vitals) != 1 {
		t.Fatalf("expected 1 vital, got %d", len(vitals))
	}
}

func TestGetVitalsByAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateVitalRequest{
		PatientID:     1,
		DoctorID:      1,
		AppointmentID: 20,
	})

	vitals, err := service.GetByAppointmentID(20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(vitals) != 1 {
		t.Fatalf("expected 1 vital, got %d", len(vitals))
	}
}

func TestUpdateVital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVitalRequest{
		PatientID:   1,
		DoctorID:    1,
		Temperature: 36.5,
	})

	temp := 38.2
	notes := "Patient has a mild fever."

	updated, err := service.Update(created.ID, UpdateVitalRequest{
		Temperature: &temp,
		Notes:       &notes,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Temperature != 38.2 {
		t.Fatal("temperature was not updated")
	}

	if updated.Notes != notes {
		t.Fatal("notes were not updated")
	}
}

func TestDeleteVital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVitalRequest{
		PatientID: 1,
		DoctorID:  1,
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrVitalNotFound {
		t.Fatalf("expected %v, got %v", ErrVitalNotFound, err)
	}
}
