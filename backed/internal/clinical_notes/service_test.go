package clinical_notes

import "testing"

func TestCreateClinicalNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		DiagnosisID:   1,
		Subject:       "Initial Consultation",
		Note:          "Patient presents with persistent cough.",
		Assessment:    "Possible viral infection.",
		Plan:          "Rest, hydration, and follow-up if symptoms persist.",
	}

	note, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if note.ID == 0 {
		t.Fatal("expected clinical note ID to be assigned")
	}

	if note.Subject != req.Subject {
		t.Fatalf("expected subject %q, got %q", req.Subject, note.Subject)
	}

	if !note.IsConfidential {
		t.Fatal("expected note to be confidential by default")
	}
}

func TestCreateClinicalNoteValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateClinicalNoteRequest{
		{},
		{
			DoctorID:  1,
			PatientID: 1,
			Subject:   "Subject",
			Note:      "Note",
		},
		{
			AppointmentID: 1,
			PatientID:     1,
			Subject:       "Subject",
			Note:          "Note",
		},
		{
			AppointmentID: 1,
			DoctorID:      1,
			Subject:       "Subject",
			Note:          "Note",
		},
		{
			AppointmentID: 1,
			DoctorID:      1,
			PatientID:     1,
			Note:          "Note",
		},
		{
			AppointmentID: 1,
			DoctorID:      1,
			PatientID:     1,
			Subject:       "Subject",
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidClinicalNote {
			t.Fatalf("expected %v, got %v", ErrInvalidClinicalNote, err)
		}
	}
}

func TestGetClinicalNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Subject:       "Consultation",
		Note:          "Clinical findings.",
	})

	note, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if note.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, note.ID)
	}
}

func TestGetByAppointmentID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateClinicalNoteRequest{
		AppointmentID: 10,
		DoctorID:      1,
		PatientID:     1,
		Subject:       "Consultation",
		Note:          "Appointment note.",
	})

	notes, err := service.GetByAppointmentID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != 1 {
		t.Fatalf("expected 1 clinical note, got %d", len(notes))
	}
}

func TestGetByDoctorID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     1,
		Subject:       "Consultation",
		Note:          "Doctor's note.",
	})

	notes, err := service.GetByDoctorID(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != 1 {
		t.Fatalf("expected 1 clinical note, got %d", len(notes))
	}
}

func TestGetByPatientID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     5,
		Subject:       "Consultation",
		Note:          "Patient note.",
	})

	notes, err := service.GetByPatientID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != 1 {
		t.Fatalf("expected 1 clinical note, got %d", len(notes))
	}
}

func TestGetByDiagnosisID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		DiagnosisID:   8,
		Subject:       "Diagnosis",
		Note:          "Linked diagnosis.",
	})

	notes, err := service.GetByDiagnosisID(8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != 1 {
		t.Fatalf("expected 1 clinical note, got %d", len(notes))
	}
}

func TestGetConfidential(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	confidential := false

	_, _ = service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Subject:       "Private",
		Note:          "Confidential note.",
	})

	_, _ = service.Create(CreateClinicalNoteRequest{
		AppointmentID:  2,
		DoctorID:       1,
		PatientID:      2,
		Subject:        "Public",
		Note:           "Visible note.",
		IsConfidential: &confidential,
	})

	notes, err := service.GetConfidential()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(notes) != 1 {
		t.Fatalf("expected 1 confidential note, got %d", len(notes))
	}
}

func TestUpdateClinicalNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Subject:       "Initial",
		Note:          "Initial note.",
	})

	confidential := false

	updated, err := service.Update(created.ID, UpdateClinicalNoteRequest{
		Subject:        "Updated",
		Assessment:     "Improving",
		Plan:           "Continue medication",
		IsConfidential: &confidential,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Subject != "Updated" {
		t.Fatal("subject was not updated")
	}

	if updated.Assessment != "Improving" {
		t.Fatal("assessment was not updated")
	}

	if updated.IsConfidential {
		t.Fatal("expected note to be non-confidential")
	}
}

func TestDeleteClinicalNote(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateClinicalNoteRequest{
		AppointmentID: 1,
		DoctorID:      1,
		PatientID:     1,
		Subject:       "Delete Test",
		Note:          "Delete me.",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrClinicalNoteNotFound {
		t.Fatalf("expected %v, got %v", ErrClinicalNoteNotFound, err)
	}
}
