package video_consultations

import (
	"testing"
	"time"
)

func TestCreateVideoConsultation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateVideoConsultationRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now().Add(time.Hour),
		Notes:         "Initial consultation",
	}

	consultation, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if consultation.ID == 0 {
		t.Fatal("expected consultation ID to be assigned")
	}

	if consultation.Status != StatusScheduled {
		t.Fatalf("expected status %s, got %s", StatusScheduled, consultation.Status)
	}

	if consultation.RoomID == "" {
		t.Fatal("expected room ID to be generated")
	}
}

func TestCreateVideoConsultationValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, err := service.Create(CreateVideoConsultationRequest{})
	if err != ErrInvalidConsultation {
		t.Fatalf("expected %v, got %v", ErrInvalidConsultation, err)
	}
}

func TestDuplicateAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateVideoConsultationRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now(),
	}

	_, _ = service.Create(req)

	_, err := service.Create(req)
	if err != ErrConsultationExists {
		t.Fatalf("expected %v, got %v", ErrConsultationExists, err)
	}
}

func TestGetVideoConsultation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVideoConsultationRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now(),
	})

	consultation, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if consultation.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, consultation.ID)
	}
}

func TestGetByAppointmentID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVideoConsultationRequest{
		AppointmentID: 10,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now(),
	})

	consultation, err := service.GetByAppointmentID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if consultation.ID != created.ID {
		t.Fatal("unexpected consultation returned")
	}
}

func TestGetDoctorConsultations(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(CreateVideoConsultationRequest{
			AppointmentID: uint(i + 1),
			DoctorID:      5,
			PatientID:     uint(i + 10),
			ScheduledAt:   time.Now(),
		})
	}

	list, err := service.GetByDoctorID(5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 3 {
		t.Fatalf("expected 3 consultations, got %d", len(list))
	}
}

func TestGetPatientConsultations(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 2; i++ {
		_, _ = service.Create(CreateVideoConsultationRequest{
			AppointmentID: uint(i + 1),
			DoctorID:      1,
			PatientID:     20,
			ScheduledAt:   time.Now(),
		})
	}

	list, err := service.GetByPatientID(20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 2 {
		t.Fatalf("expected 2 consultations, got %d", len(list))
	}
}

func TestUpdateVideoConsultation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVideoConsultationRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now(),
	})

	status := StatusOngoing
	notes := "Consultation started"

	updated, err := service.Update(created.ID, UpdateVideoConsultationRequest{
		Status: &status,
		Notes:  &notes,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Status != StatusOngoing {
		t.Fatal("status not updated")
	}

	if updated.StartedAt == nil {
		t.Fatal("expected StartedAt to be set")
	}

	if updated.Notes != notes {
		t.Fatal("notes not updated")
	}
}

func TestCompleteVideoConsultation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVideoConsultationRequest{
		AppointmentID: 2,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now(),
	})

	status := StatusCompleted

	updated, err := service.Update(created.ID, UpdateVideoConsultationRequest{
		Status: &status,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.EndedAt == nil {
		t.Fatal("expected EndedAt to be set")
	}
}

func TestDeleteVideoConsultation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateVideoConsultationRequest{
		AppointmentID: 1,
		DoctorID:      2,
		PatientID:     3,
		ScheduledAt:   time.Now(),
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrConsultationNotFound {
		t.Fatalf("expected %v, got %v", ErrConsultationNotFound, err)
	}
}

func TestGetAllVideoConsultations(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 5; i++ {
		_, _ = service.Create(CreateVideoConsultationRequest{
			AppointmentID: uint(i + 1),
			DoctorID:      1,
			PatientID:     uint(i + 10),
			ScheduledAt:   time.Now(),
		})
	}

	list, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 5 {
		t.Fatalf("expected 5 consultations, got %d", len(list))
	}
}
