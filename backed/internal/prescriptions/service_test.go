package prescriptions

import "testing"

func TestCreatePrescription(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreatePrescriptionRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 3,
		Diagnosis:     "Malaria",
		Notes:         "Patient should rest and stay hydrated.",
		Items: []PrescriptionItemRequest{
			{
				MedicationName: "Coartem",
				Dosage:         "20/120mg",
				Frequency:      "Twice daily",
				Duration:       "3 days",
				Instructions:   "After meals",
				Quantity:       6,
				Refills:        0,
			},
		},
	}

	prescription, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if prescription.ID == 0 {
		t.Fatal("expected prescription ID to be assigned")
	}

	if prescription.Status != StatusActive {
		t.Fatalf("expected status %s, got %s", StatusActive, prescription.Status)
	}

	if len(prescription.Items) != 1 {
		t.Fatalf("expected 1 prescription item, got %d", len(prescription.Items))
	}
}

func TestCreatePrescriptionValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreatePrescriptionRequest{
		{},
		{
			DoctorID:      1,
			PatientID:     1,
			AppointmentID: 1,
			Diagnosis:     "",
			Items: []PrescriptionItemRequest{
				{
					MedicationName: "Drug",
					Dosage:         "500mg",
					Frequency:      "Daily",
					Duration:       "7 days",
					Quantity:       1,
				},
			},
		},
		{
			DoctorID:      1,
			PatientID:     1,
			AppointmentID: 1,
			Diagnosis:     "Diagnosis",
			Items:         []PrescriptionItemRequest{},
		},
	}

	for _, req := range tests {
		_, err := service.Create(req)
		if err != ErrInvalidPrescription {
			t.Fatalf("expected %v, got %v", ErrInvalidPrescription, err)
		}
	}
}

func TestGetPrescription(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	prescription, _ := service.Create(CreatePrescriptionRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 3,
		Diagnosis:     "Flu",
		Items: []PrescriptionItemRequest{
			{
				MedicationName: "Paracetamol",
				Dosage:         "500mg",
				Frequency:      "Three times daily",
				Duration:       "5 days",
				Quantity:       15,
			},
		},
	})

	found, err := service.GetByID(prescription.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != prescription.ID {
		t.Fatalf("expected ID %d, got %d", prescription.ID, found.ID)
	}
}

func TestGetAllPrescriptions(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(CreatePrescriptionRequest{
			DoctorID:      1,
			PatientID:     uint(i + 1),
			AppointmentID: uint(i + 1),
			Diagnosis:     "Diagnosis",
			Items: []PrescriptionItemRequest{
				{
					MedicationName: "Drug",
					Dosage:         "100mg",
					Frequency:      "Daily",
					Duration:       "7 days",
					Quantity:       7,
				},
			},
		})
	}

	prescriptions, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(prescriptions) != 3 {
		t.Fatalf("expected 3 prescriptions, got %d", len(prescriptions))
	}
}

func TestGetByDoctorID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreatePrescriptionRequest{
		DoctorID:      10,
		PatientID:     1,
		AppointmentID: 1,
		Diagnosis:     "Diagnosis",
		Items: []PrescriptionItemRequest{
			{
				MedicationName: "Drug",
				Dosage:         "100mg",
				Frequency:      "Daily",
				Duration:       "7 days",
				Quantity:       7,
			},
		},
	})

	prescriptions, err := service.GetByDoctorID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(prescriptions) != 1 {
		t.Fatalf("expected 1 prescription, got %d", len(prescriptions))
	}
}

func TestGetByPatientID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreatePrescriptionRequest{
		DoctorID:      1,
		PatientID:     20,
		AppointmentID: 1,
		Diagnosis:     "Diagnosis",
		Items: []PrescriptionItemRequest{
			{
				MedicationName: "Drug",
				Dosage:         "100mg",
				Frequency:      "Daily",
				Duration:       "7 days",
				Quantity:       7,
			},
		},
	})

	prescriptions, err := service.GetByPatientID(20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(prescriptions) != 1 {
		t.Fatalf("expected 1 prescription, got %d", len(prescriptions))
	}
}

func TestUpdatePrescription(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	prescription, _ := service.Create(CreatePrescriptionRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 3,
		Diagnosis:     "Initial Diagnosis",
		Items: []PrescriptionItemRequest{
			{
				MedicationName: "Drug A",
				Dosage:         "100mg",
				Frequency:      "Daily",
				Duration:       "7 days",
				Quantity:       7,
			},
		},
	})

	updated, err := service.Update(prescription.ID, UpdatePrescriptionRequest{
		Diagnosis: "Updated Diagnosis",
		Status:    StatusCompleted,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Diagnosis != "Updated Diagnosis" {
		t.Fatal("diagnosis was not updated")
	}

	if updated.Status != StatusCompleted {
		t.Fatal("status was not updated")
	}
}

func TestDeletePrescription(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	prescription, _ := service.Create(CreatePrescriptionRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 3,
		Diagnosis:     "Diagnosis",
		Items: []PrescriptionItemRequest{
			{
				MedicationName: "Drug",
				Dosage:         "100mg",
				Frequency:      "Daily",
				Duration:       "7 days",
				Quantity:       7,
			},
		},
	})

	if err := service.Delete(prescription.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(prescription.ID)
	if err != ErrPrescriptionNotFound {
		t.Fatalf("expected %v, got %v", ErrPrescriptionNotFound, err)
	}
}
