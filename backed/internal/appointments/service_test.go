package appointments

import (
	"testing"
	"time"
)

func TestCreateAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateAppointmentRequest{
		DoctorID:        2,
		AppointmentTime: time.Now().Add(24 * time.Hour),
		DurationMinutes: 45,
		Reason:          "General consultation",
	}

	appointment, err := service.CreateAppointment(1, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if appointment.ID == 0 {
		t.Errorf("expected appointment ID to be set")
	}

	if appointment.PatientID != 1 {
		t.Errorf("expected patient ID 1, got %d", appointment.PatientID)
	}

	if appointment.DoctorID != 2 {
		t.Errorf("expected doctor ID 2, got %d", appointment.DoctorID)
	}

	if appointment.Status != StatusPending {
		t.Errorf("expected status %s, got %s", StatusPending, appointment.Status)
	}
}

func TestCreateAppointmentValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []struct {
		name string
		req  CreateAppointmentRequest
	}{
		{
			name: "missing doctor",
			req: CreateAppointmentRequest{
				AppointmentTime: time.Now(),
				Reason:          "Consultation",
			},
		},
		{
			name: "missing reason",
			req: CreateAppointmentRequest{
				DoctorID:        2,
				AppointmentTime: time.Now(),
			},
		},
	}

	for _, tt := range tests {
		_, err := service.CreateAppointment(1, tt.req)
		if err == nil {
			t.Errorf("%s: expected validation error", tt.name)
		}
	}
}

func TestSelfAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateAppointmentRequest{
		DoctorID:        1,
		AppointmentTime: time.Now(),
		Reason:          "Consultation",
	}

	_, err := service.CreateAppointment(1, req)
	if err == nil {
		t.Fatal("expected self-appointment error")
	}
}

func TestGetAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateAppointmentRequest{
		DoctorID:        2,
		AppointmentTime: time.Now(),
		Reason:          "Consultation",
	}

	created, _ := service.CreateAppointment(1, req)

	appointment, err := service.GetAppointment(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if appointment.ID != created.ID {
		t.Errorf("expected appointment ID %d, got %d", created.ID, appointment.ID)
	}
}

func TestUpdateAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateAppointmentRequest{
		DoctorID:        2,
		AppointmentTime: time.Now(),
		Reason:          "Consultation",
	}

	created, _ := service.CreateAppointment(1, req)

	status := StatusConfirmed
	notes := "Doctor confirmed appointment"

	updateReq := UpdateAppointmentRequest{
		Status: &status,
		Notes:  &notes,
	}

	updated, err := service.UpdateAppointment(created.ID, updateReq)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if updated.Status != StatusConfirmed {
		t.Errorf("expected status %s, got %s", StatusConfirmed, updated.Status)
	}

	if updated.Notes != notes {
		t.Errorf("expected notes to be updated")
	}
}

func TestGetDoctorAppointments(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	service.CreateAppointment(1, CreateAppointmentRequest{
		DoctorID:        5,
		AppointmentTime: time.Now(),
		Reason:          "Consultation",
	})

	service.CreateAppointment(2, CreateAppointmentRequest{
		DoctorID:        5,
		AppointmentTime: time.Now(),
		Reason:          "Follow-up",
	})

	appointments, err := service.GetDoctorAppointments(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(appointments) != 2 {
		t.Errorf("expected 2 appointments, got %d", len(appointments))
	}
}

func TestGetPatientAppointments(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, err := service.CreateAppointment(10, CreateAppointmentRequest{
		DoctorID:        2,
		AppointmentTime: time.Now(),
		Reason:          "Consultation",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = service.CreateAppointment(10, CreateAppointmentRequest{
		DoctorID:        3,
		AppointmentTime: time.Now(),
		Reason:          "Review",
	})

	appointments, err := service.GetPatientAppointments(10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(appointments) != 2 {
		t.Errorf("expected 2 appointments, got %d", len(appointments))
	}
}

func TestDeleteAppointment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	// 1. Capture the created appointment object to get its real ID
	created, err := service.CreateAppointment(1, CreateAppointmentRequest{
		DoctorID:        2,
		AppointmentTime: time.Now(),
		Reason:          "Consultation",
	})
	if err != nil {
		t.Fatalf("expected no creation error, got %v", err)
	}

	// 2. Execute the actual deletion operation
	if err := service.DeleteAppointment(created.ID); err != nil {
		t.Fatalf("expected no error during deletion, got %v", err)
	}

	// 3. Query the record again to verify it is gone
	_, err = service.GetAppointment(created.ID)
	if err == nil {
		t.Fatal("expected appointment to be deleted, but it was found")
	}
}
