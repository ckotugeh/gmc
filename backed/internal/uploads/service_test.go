package uploads

import "testing"

func TestCreateUpload(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	upload := &Upload{
		UserID:       1,
		FileName:     "report_123.pdf",
		OriginalName: "medical_report.pdf",
		FileType:     "application/pdf",
		FileSize:     2048,
		FilePath:     "/uploads/report_123.pdf",
		Description:  "Medical report",
		IsPublic:     false,
	}

	created, err := service.Create(upload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if created.ID == 0 {
		t.Fatal("expected upload ID to be assigned")
	}
}

func TestCreateUploadValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	upload := &Upload{}

	_, err := service.Create(upload)
	if err != ErrInvalidUpload {
		t.Fatalf("expected %v, got %v", ErrInvalidUpload, err)
	}
}

func TestGetUpload(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, err := service.Create(&Upload{
		UserID:       1,
		FileName:     "image.png",
		OriginalName: "image.png",
		FileType:     "image/png",
		FileSize:     1000,
		FilePath:     "/uploads/image.png",
	})

	upload, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if upload.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, upload.ID)
	}
}

func TestGetByUserID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	service.Create(&Upload{
		UserID:       1,
		FileName:     "a.pdf",
		OriginalName: "a.pdf",
		FileType:     "application/pdf",
		FileSize:     100,
		FilePath:     "/uploads/a.pdf",
	})

	service.Create(&Upload{
		UserID:       1,
		FileName:     "b.pdf",
		OriginalName: "b.pdf",
		FileType:     "application/pdf",
		FileSize:     100,
		FilePath:     "/uploads/b.pdf",
	})

	service.Create(&Upload{
		UserID:       2,
		FileName:     "c.pdf",
		OriginalName: "c.pdf",
		FileType:     "application/pdf",
		FileSize:     100,
		FilePath:     "/uploads/c.pdf",
	})

	uploads, err := service.GetByUserID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(uploads) != 2 {
		t.Fatalf("expected 2 uploads, got %d", len(uploads))
	}
}

func TestUpdateUpload(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, err := service.Create(&Upload{
		UserID:       1,
		FileName:     "report.pdf",
		OriginalName: "report.pdf",
		FileType:     "application/pdf",
		FileSize:     500,
		FilePath:     "/uploads/report.pdf",
	})

	public := true

	updated, err := service.Update(created.ID, UpdateUploadRequest{
		Description: "Updated report",
		IsPublic:    &public,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Description != "Updated report" {
		t.Errorf("expected updated description")
	}

	if !updated.IsPublic {
		t.Errorf("expected upload to be public")
	}
}

func TestDeleteUpload(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, err := service.Create(&Upload{
		UserID:       1,
		FileName:     "delete.pdf",
		OriginalName: "delete.pdf",
		FileType:     "application/pdf",
		FileSize:     100,
		FilePath:     "/uploads/delete.pdf",
	})
	if err != nil {
		t.Fatalf("unexpected setup error: %v", err)
	}

	// 1. Perform deletion step
	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error during deletion: %v", err)
	}

	// 2. Query the record again to verify it returns ErrUploadNotFound
	_, err = service.GetByID(created.ID)
	if err != ErrUploadNotFound {
		t.Fatalf("expected %v, got %v", ErrUploadNotFound, err)
	}
}

func TestGetByAppointmentID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	appointmentID := uint(10)

	service.Create(&Upload{
		UserID:        1,
		AppointmentID: &appointmentID,
		FileName:      "appointment.pdf",
		OriginalName:  "appointment.pdf",
		FileType:      "application/pdf",
		FileSize:      500,
		FilePath:      "/uploads/appointment.pdf",
	})

	uploads, err := service.GetByAppointmentID(appointmentID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(uploads) != 1 {
		t.Fatalf("expected 1 upload, got %d", len(uploads))
	}
}

func TestGetByMedicalRecordID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	recordID := uint(20)

	service.Create(&Upload{
		UserID:          1,
		MedicalRecordID: &recordID,
		FileName:        "record.pdf",
		OriginalName:    "record.pdf",
		FileType:        "application/pdf",
		FileSize:        500,
		FilePath:        "/uploads/record.pdf",
	})

	uploads, err := service.GetByMedicalRecordID(recordID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(uploads) != 1 {
		t.Fatalf("expected 1 upload, got %d", len(uploads))
	}
}

func TestGetByHospitalID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	hospitalID := uint(30)

	service.Create(&Upload{
		UserID:       1,
		HospitalID:   &hospitalID,
		FileName:     "hospital.pdf",
		OriginalName: "hospital.pdf",
		FileType:     "application/pdf",
		FileSize:     500,
		FilePath:     "/uploads/hospital.pdf",
	})

	uploads, err := service.GetByHospitalID(hospitalID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(uploads) != 1 {
		t.Fatalf("expected 1 upload, got %d", len(uploads))
	}
}
