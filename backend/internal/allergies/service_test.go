package allergies

import "testing"

func TestCreateAllergy(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Penicillin",
		Type:      "Drug",
		Severity:  "Severe",
		Reaction:  "Anaphylaxis",
		Status:    "Active",
		Notes:     "Patient must avoid all penicillin-based antibiotics.",
	}

	allergy, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if allergy.ID == 0 {
		t.Fatal("expected allergy ID to be assigned")
	}

	if allergy.Allergen != "Penicillin" {
		t.Fatalf("expected allergen Penicillin, got %s", allergy.Allergen)
	}
}

func TestCreateAllergyValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateAllergyRequest{
		{},
		{
			DoctorID: 1,
			Allergen: "Peanuts",
			Type:     "Food",
			Severity: "Moderate",
		},
		{
			PatientID: 1,
			Allergen:  "Peanuts",
			Type:      "Food",
			Severity:  "Moderate",
		},
		{
			PatientID: 1,
			DoctorID:  1,
			Type:      "Food",
			Severity:  "Moderate",
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidAllergy {
			t.Fatalf("expected %v, got %v", ErrInvalidAllergy, err)
		}
	}
}

func TestGetAllergyByID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Dust",
		Type:      "Environmental",
		Severity:  "Mild",
	})

	allergy, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if allergy.ID != created.ID {
		t.Fatal("returned wrong allergy")
	}
}

func TestGetAllAllergies(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Dust",
		Type:      "Environmental",
		Severity:  "Mild",
	})

	_, _ = service.Create(CreateAllergyRequest{
		PatientID: 2,
		DoctorID:  1,
		Allergen:  "Peanuts",
		Type:      "Food",
		Severity:  "Severe",
	})

	allergies, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(allergies) != 2 {
		t.Fatalf("expected 2 allergies, got %d", len(allergies))
	}
}

func TestGetAllergiesByPatient(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAllergyRequest{
		PatientID: 10,
		DoctorID:  1,
		Allergen:  "Dust",
		Type:      "Environmental",
		Severity:  "Mild",
	})

	allergies, err := service.GetByPatientID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(allergies) != 1 {
		t.Fatalf("expected 1 allergy, got %d", len(allergies))
	}
}

func TestGetAllergiesByDoctor(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  20,
		Allergen:  "Dust",
		Type:      "Environmental",
		Severity:  "Mild",
	})

	allergies, err := service.GetByDoctorID(20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(allergies) != 1 {
		t.Fatalf("expected 1 allergy, got %d", len(allergies))
	}
}

func TestGetAllergiesBySeverity(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Peanuts",
		Type:      "Food",
		Severity:  "Severe",
	})

	allergies, err := service.GetBySeverity("Severe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(allergies) != 1 {
		t.Fatalf("expected 1 allergy, got %d", len(allergies))
	}
}

func TestGetActiveAllergies(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Latex",
		Type:      "Other",
		Severity:  "Moderate",
		Status:    "Active",
	})

	allergies, err := service.GetActive()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(allergies) != 1 {
		t.Fatalf("expected 1 allergy, got %d", len(allergies))
	}
}

func TestUpdateAllergy(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Dust",
		Type:      "Environmental",
		Severity:  "Mild",
	})

	severity := "Severe"
	notes := "Patient now experiences breathing difficulties."

	updated, err := service.Update(created.ID, UpdateAllergyRequest{
		Severity: &severity,
		Notes:    &notes,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Severity != "Severe" {
		t.Fatal("severity was not updated")
	}

	if updated.Notes != notes {
		t.Fatal("notes were not updated")
	}
}

func TestDeleteAllergy(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateAllergyRequest{
		PatientID: 1,
		DoctorID:  1,
		Allergen:  "Dust",
		Type:      "Environmental",
		Severity:  "Mild",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrAllergyNotFound {
		t.Fatalf("expected %v, got %v", ErrAllergyNotFound, err)
	}
}
