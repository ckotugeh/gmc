package lab_results

import "testing"

func TestCreateLabResult(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateLabResultRequest{
		LabRequestID:   1,
		PatientID:      1,
		DoctorID:       1,
		TestName:       "Complete Blood Count",
		Result:         "Hemoglobin: 13.8 g/dL",
		ReferenceRange: "12.0 - 16.0 g/dL",
		Units:          "g/dL",
		Interpretation: "Normal",
		Status:         "Completed",
		Remarks:        "No abnormalities detected.",
	}

	result, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.ID == 0 {
		t.Fatal("expected result ID to be assigned")
	}

	if result.TestName != req.TestName {
		t.Fatalf("expected %s, got %s", req.TestName, result.TestName)
	}
}

func TestCreateLabResultValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateLabResultRequest{
		{},
		{
			PatientID: 1,
			DoctorID:  1,
			TestName:  "CBC",
			Result:    "Normal",
		},
		{
			LabRequestID: 1,
			DoctorID:     1,
			TestName:     "CBC",
			Result:       "Normal",
		},
		{
			LabRequestID: 1,
			PatientID:    1,
			TestName:     "CBC",
			Result:       "Normal",
		},
		{
			LabRequestID: 1,
			PatientID:    1,
			DoctorID:     1,
			Result:       "Normal",
		},
		{
			LabRequestID: 1,
			PatientID:    1,
			DoctorID:     1,
			TestName:     "CBC",
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidLabResult {
			t.Fatalf("expected %v, got %v", ErrInvalidLabResult, err)
		}
	}
}

func TestGetLabResultByID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    1,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
	})

	result, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.ID != created.ID {
		t.Fatal("returned wrong lab result")
	}
}

func TestGetAllLabResults(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    1,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
	})

	_, _ = service.Create(CreateLabResultRequest{
		LabRequestID: 2,
		PatientID:    2,
		DoctorID:     2,
		TestName:     "Liver Function Test",
		Result:       "Elevated ALT",
	})

	results, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
}

func TestGetLabResultsByLabRequest(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabResultRequest{
		LabRequestID: 5,
		PatientID:    1,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
	})

	results, err := service.GetByLabRequestID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestGetLabResultsByPatient(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    7,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
	})

	results, err := service.GetByPatientID(7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestGetLabResultsByDoctor(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    1,
		DoctorID:     10,
		TestName:     "CBC",
		Result:       "Normal",
	})

	results, err := service.GetByDoctorID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestGetLabResultsByStatus(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    1,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
		Status:       "Verified",
	})

	results, err := service.GetByStatus("Verified")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestUpdateLabResult(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    1,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
	})

	status := "Verified"
	remarks := "Reviewed by pathologist."

	updated, err := service.Update(created.ID, UpdateLabResultRequest{
		Status:  &status,
		Remarks: &remarks,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Status != "Verified" {
		t.Fatal("status was not updated")
	}

	if updated.Remarks != remarks {
		t.Fatal("remarks were not updated")
	}
}

func TestDeleteLabResult(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateLabResultRequest{
		LabRequestID: 1,
		PatientID:    1,
		DoctorID:     1,
		TestName:     "CBC",
		Result:       "Normal",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrLabResultNotFound {
		t.Fatalf("expected %v, got %v", ErrLabResultNotFound, err)
	}
}
