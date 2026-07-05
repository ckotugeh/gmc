package availability

import (
	"testing"
	"time"
)

func TestCreateAvailability(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	date := time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC)

	req := CreateAvailabilityRequest{
		DoctorID:   1,
		ScheduleID: 1,
		Date:       date,
		StartTime:  "09:00",
		EndTime:    "09:30",
	}

	slot, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if slot.ID == 0 {
		t.Fatal("expected slot ID to be assigned")
	}

	if slot.Status != SlotAvailable {
		t.Fatalf("expected status %s, got %s", SlotAvailable, slot.Status)
	}
}

func TestCreateAvailabilityValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateAvailabilityRequest{
		{},
		{
			DoctorID:   0,
			ScheduleID: 1,
			Date:       time.Now(),
			StartTime:  "09:00",
			EndTime:    "09:30",
		},
		{
			DoctorID:   1,
			ScheduleID: 0,
			Date:       time.Now(),
			StartTime:  "09:00",
			EndTime:    "09:30",
		},
		{
			DoctorID:   1,
			ScheduleID: 1,
			StartTime:  "09:00",
			EndTime:    "09:30",
		},
		{
			DoctorID:   1,
			ScheduleID: 1,
			Date:       time.Now(),
			StartTime:  "",
			EndTime:    "09:30",
		},
		{
			DoctorID:   1,
			ScheduleID: 1,
			Date:       time.Now(),
			StartTime:  "09:00",
			EndTime:    "",
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidAvailability {
			t.Fatalf("expected %v, got %v", ErrInvalidAvailability, err)
		}
	}
}

func TestGetAvailability(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	slot, _ := service.Create(CreateAvailabilityRequest{
		DoctorID:   1,
		ScheduleID: 1,
		Date:       time.Now(),
		StartTime:  "10:00",
		EndTime:    "10:30",
	})

	found, err := service.GetByID(slot.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != slot.ID {
		t.Fatalf("expected ID %d, got %d", slot.ID, found.ID)
	}
}

func TestGetAllAvailability(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(CreateAvailabilityRequest{
			DoctorID:   uint(i + 1),
			ScheduleID: 1,
			Date:       time.Now(),
			StartTime:  "09:00",
			EndTime:    "09:30",
		})
	}

	slots, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(slots) != 3 {
		t.Fatalf("expected 3 slots, got %d", len(slots))
	}
}

func TestGetByDoctorID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAvailabilityRequest{
		DoctorID:   10,
		ScheduleID: 1,
		Date:       time.Now(),
		StartTime:  "11:00",
		EndTime:    "11:30",
	})

	slots, err := service.GetByDoctorID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(slots) != 1 {
		t.Fatalf("expected 1 slot, got %d", len(slots))
	}
}

func TestGetByScheduleID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAvailabilityRequest{
		DoctorID:   1,
		ScheduleID: 20,
		Date:       time.Now(),
		StartTime:  "13:00",
		EndTime:    "13:30",
	})

	slots, err := service.GetByScheduleID(20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(slots) != 1 {
		t.Fatalf("expected 1 slot, got %d", len(slots))
	}
}

func TestGetByDoctorAndDate(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	date := time.Date(2026, 1, 6, 0, 0, 0, 0, time.UTC)

	_, _ = service.Create(CreateAvailabilityRequest{
		DoctorID:   7,
		ScheduleID: 1,
		Date:       date,
		StartTime:  "15:00",
		EndTime:    "15:30",
	})

	slots, err := service.GetByDoctorAndDate(7, date)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(slots) != 1 {
		t.Fatalf("expected 1 slot, got %d", len(slots))
	}
}

func TestUpdateAvailability(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	slot, _ := service.Create(CreateAvailabilityRequest{
		DoctorID:   1,
		ScheduleID: 1,
		Date:       time.Now(),
		StartTime:  "09:00",
		EndTime:    "09:30",
	})

	updated, err := service.Update(slot.ID, UpdateAvailabilityRequest{
		StartTime: "10:00",
		EndTime:   "10:30",
		Status:    SlotBooked,
		Notes:     "Reserved",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.StartTime != "10:00" {
		t.Fatal("start time was not updated")
	}

	if updated.EndTime != "10:30" {
		t.Fatal("end time was not updated")
	}

	if updated.Status != SlotBooked {
		t.Fatal("status was not updated")
	}

	if updated.Notes != "Reserved" {
		t.Fatal("notes were not updated")
	}
}

func TestDeleteAvailability(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	slot, _ := service.Create(CreateAvailabilityRequest{
		DoctorID:   1,
		ScheduleID: 1,
		Date:       time.Now(),
		StartTime:  "09:00",
		EndTime:    "09:30",
	})

	if err := service.Delete(slot.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(slot.ID)
	if err != ErrAvailabilityNotFound {
		t.Fatalf("expected %v, got %v", ErrAvailabilityNotFound, err)
	}
}
