package lab_requests

import "testing"

func TestCreateLabRequest(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateLabRequestRequest{
		PatientID:     1,
		DoctorID:      1,
		AppointmentID: 1,
		TestName:      "Complete Blood Count",
		Category:      "Hematology",
		Priority:      "Routine",
		ClinicalNotes: "Patient complains of fatigue.",
		Reason:        "Rule out anemia",
		Status:        "Pending",
	}

	request, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if request.ID == 0 {
		t.Fatal("expected request ID to be assigned")
	}

	if request.TestName != req.TestName {
		t.Fatalf("expected %s, got %s", req.TestName, request.TestName)
	}
}

func TestCreateLabRequestValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateLabRequestRequest{
		{},
		{
			DoctorID: 1,
			TestName: "CBC",
		},
		{
			PatientID: 1,
			TestName:  "CBC",
		},
		{
			PatientID: 1,
			DoctorID:  1,
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidLabRequest {
			t.Fatalf("expected %v, got %v", ErrInvalidLabRequest, err)
		}
	}
}

func TestGetLabRequestByID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateLabRequestRequest{
		PatientID: 1,
		DoctorID:  1,
		TestName:  "CBC",
	})

	request, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if request.ID != created.ID {
		t.Fatal("returned wrong lab request")
	}
}

func TestGetAllLabRequests(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabRequestRequest{
		PatientID: 1,
		DoctorID:  1,
		TestName:  "CBC",
	})

	_, _ = service.Create(CreateLabRequestRequest{
		PatientID: 2,
		DoctorID:  2,
		TestName:  "Liver Function Test",
	})

	requests, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(requests) != 2 {
		t.Fatalf("expected 2 requests, got %d", len(requests))
	}
}

func TestGetLabRequestsByPatient(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabRequestRequest{
		PatientID: 5,
		DoctorID:  1,
		TestName:  "CBC",
	})

	requests, err := service.GetByPatientID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestGetLabRequestsByDoctor(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabRequestRequest{
		PatientID: 1,
		DoctorID:  9,
		TestName:  "Urinalysis",
	})

	requests, err := service.GetByDoctorID(9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestGetLabRequestsByAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabRequestRequest{
		PatientID:     1,
		DoctorID:      1,
		AppointmentID: 100,
		TestName:      "Blood Glucose",
	})

	requests, err := service.GetByAppointmentID(100)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestGetLabRequestsByStatus(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabRequestRequest{
		PatientID: 1,
		DoctorID:  1,
		TestName:  "CBC",
		Status:    "Pending",
	})

	requests, err := service.GetByStatus("Pending")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(requests) != 1 {
		t.Fatalf("expected 1 request, got %d", len(requests))
	}
}

func TestUpdateLabRequest(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateLabRequestRequest{
		PatientID: 1,
		DoctorID:  1,
		TestName:  "CBC",
	})

	status := "Completed"
	notes := "Results available."

	updated, err := service.Update(created.ID, UpdateLabRequestRequest{
		Status:        &status,
		ClinicalNotes: &notes,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Status != "Completed" {
		t.Fatal("status was not updated")
	}

	if updated.ClinicalNotes != notes {
		t.Fatal("clinical notes were not updated")
	}
}

func TestDeleteLabRequest(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateLabRequestRequest{
		PatientID: 1,
		DoctorID:  1,
		TestName:  "CBC",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrLabRequestNotFound {
		t.Fatalf("expected %v, got %v", ErrLabRequestNotFound, err)
	}
}
