package hospitals

import "testing"

func TestCreateHospital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateHospitalRequest{
		Name:          "Nairobi Hospital",
		Email:         "info@nairobihospital.com",
		Phone:         "+254700000000",
		Address:       "Argwings Kodhek Rd",
		City:          "Nairobi",
		Country:       "Kenya",
		LicenseNumber: "NH-001",
		IsActive:      true,
	}

	hospital, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hospital.ID == 0 {
		t.Fatal("expected hospital ID to be assigned")
	}

	if hospital.Name != req.Name {
		t.Errorf("expected %q, got %q", req.Name, hospital.Name)
	}
}

func TestCreateHospitalValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateHospitalRequest{}

	_, err := service.Create(req)
	if err != ErrInvalidHospital {
		t.Fatalf("expected %v, got %v", ErrInvalidHospital, err)
	}
}

func TestCreateHospitalDuplicateEmail(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateHospitalRequest{
		Name:          "Hospital One",
		Email:         "hospital@test.com",
		LicenseNumber: "LIC-001",
	}

	if _, err := service.Create(req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req2 := CreateHospitalRequest{
		Name:          "Hospital Two",
		Email:         "hospital@test.com",
		LicenseNumber: "LIC-002",
	}

	_, err := service.Create(req2)
	if err != ErrHospitalEmailExists {
		t.Fatalf("expected %v, got %v", ErrHospitalEmailExists, err)
	}
}

func TestCreateHospitalDuplicateLicense(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateHospitalRequest{
		Name:          "Hospital One",
		Email:         "one@test.com",
		LicenseNumber: "LIC-001",
	}

	if _, err := service.Create(req); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	req2 := CreateHospitalRequest{
		Name:          "Hospital Two",
		Email:         "two@test.com",
		LicenseNumber: "LIC-001",
	}

	_, err := service.Create(req2)
	if err != ErrHospitalLicenseExists {
		t.Fatalf("expected %v, got %v", ErrHospitalLicenseExists, err)
	}
}

func TestGetHospital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateHospitalRequest{
		Name:          "Aga Khan Hospital",
		Email:         "info@agakhan.com",
		LicenseNumber: "AK-001",
	})

	hospital, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if hospital.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, hospital.ID)
	}
}

func TestGetAllHospitals(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	service.Create(CreateHospitalRequest{
		Name:          "Hospital A",
		Email:         "a@test.com",
		LicenseNumber: "A-001",
	})

	service.Create(CreateHospitalRequest{
		Name:          "Hospital B",
		Email:         "b@test.com",
		LicenseNumber: "B-001",
	})

	hospitals, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(hospitals) != 2 {
		t.Fatalf("expected 2 hospitals, got %d", len(hospitals))
	}
}

func TestUpdateHospital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateHospitalRequest{
		Name:          "Old Hospital",
		Email:         "old@test.com",
		LicenseNumber: "OLD-001",
	})

	active := false

	updated, err := service.Update(created.ID, UpdateHospitalRequest{
		Name:     "New Hospital",
		City:     "Kisumu",
		IsActive: &active,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Name != "New Hospital" {
		t.Errorf("expected updated name")
	}

	if updated.City != "Kisumu" {
		t.Errorf("expected updated city")
	}

	if updated.IsActive {
		t.Errorf("expected hospital to be inactive")
	}
}

func TestDeleteHospital(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateHospitalRequest{
		Name:          "Delete Hospital",
		Email:         "delete@test.com",
		LicenseNumber: "DEL-001",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrHospitalNotFound {
		t.Fatalf("expected %v, got %v", ErrHospitalNotFound, err)
	}
}
