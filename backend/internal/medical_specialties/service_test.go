package medical_specialties

import "testing"

func TestCreateMedicalSpecialty(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateMedicalSpecialtyRequest{
		Name:        "Cardiology",
		Code:        "CARD",
		Description: "Heart and cardiovascular diseases",
	}

	specialty, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if specialty.ID == 0 {
		t.Fatal("expected specialty ID to be assigned")
	}

	if specialty.Name != req.Name {
		t.Fatalf("expected name %q, got %q", req.Name, specialty.Name)
	}

	if specialty.Code != req.Code {
		t.Fatalf("expected code %q, got %q", req.Code, specialty.Code)
	}

	if !specialty.IsActive {
		t.Fatal("expected specialty to be active by default")
	}
}

func TestCreateMedicalSpecialtyValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateMedicalSpecialtyRequest{
		{},
		{
			Name: "",
			Code: "CARD",
		},
		{
			Name: "Cardiology",
			Code: "",
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidMedicalSpecialty {
			t.Fatalf("expected %v, got %v", ErrInvalidMedicalSpecialty, err)
		}
	}
}

func TestDuplicateMedicalSpecialtyName(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateMedicalSpecialtyRequest{
		Name: "Cardiology",
		Code: "CARD",
	}

	_, _ = service.Create(req)

	_, err := service.Create(CreateMedicalSpecialtyRequest{
		Name: "Cardiology",
		Code: "CARD2",
	})

	if err != ErrDuplicateSpecialtyName {
		t.Fatalf("expected %v, got %v", ErrDuplicateSpecialtyName, err)
	}
}

func TestDuplicateMedicalSpecialtyCode(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateMedicalSpecialtyRequest{
		Name: "Cardiology",
		Code: "CARD",
	})

	_, err := service.Create(CreateMedicalSpecialtyRequest{
		Name: "Neurology",
		Code: "CARD",
	})

	if err != ErrDuplicateSpecialtyCode {
		t.Fatalf("expected %v, got %v", ErrDuplicateSpecialtyCode, err)
	}
}

func TestGetMedicalSpecialty(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateMedicalSpecialtyRequest{
		Name: "Dermatology",
		Code: "DERM",
	})

	specialty, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if specialty.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, specialty.ID)
	}
}

func TestGetAllMedicalSpecialties(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateMedicalSpecialtyRequest{
		Name: "Cardiology",
		Code: "CARD",
	})

	_, _ = service.Create(CreateMedicalSpecialtyRequest{
		Name: "Neurology",
		Code: "NEUR",
	})

	specialties, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(specialties) != 2 {
		t.Fatalf("expected 2 specialties, got %d", len(specialties))
	}
}

func TestGetActiveMedicalSpecialties(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateMedicalSpecialtyRequest{
		Name: "Cardiology",
		Code: "CARD",
	})

	active := false

	_, _ = service.Create(CreateMedicalSpecialtyRequest{
		Name:     "Retired Specialty",
		Code:     "OLD",
		IsActive: &active,
	})

	specialties, err := service.GetActive()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(specialties) != 1 {
		t.Fatalf("expected 1 active specialty, got %d", len(specialties))
	}
}

func TestUpdateMedicalSpecialty(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateMedicalSpecialtyRequest{
		Name: "Cardiology",
		Code: "CARD",
	})

	active := false

	updated, err := service.Update(created.ID, UpdateMedicalSpecialtyRequest{
		Name:        "Interventional Cardiology",
		Description: "Updated description",
		IsActive:    &active,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Name != "Interventional Cardiology" {
		t.Fatal("name was not updated")
	}

	if updated.Description != "Updated description" {
		t.Fatal("description was not updated")
	}

	if updated.IsActive {
		t.Fatal("expected specialty to be inactive")
	}
}

func TestDeleteMedicalSpecialty(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateMedicalSpecialtyRequest{
		Name: "Radiology",
		Code: "RAD",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrMedicalSpecialtyNotFound {
		t.Fatalf("expected %v, got %v", ErrMedicalSpecialtyNotFound, err)
	}
}
