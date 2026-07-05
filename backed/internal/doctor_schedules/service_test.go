package doctor_schedules

import "testing"

func TestCreateSchedule(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateDoctorScheduleRequest{
		DoctorID:             1,
		Day:                  Monday,
		StartTime:            "08:00",
		EndTime:              "17:00",
		ConsultationDuration: 30,
		MaxPatients:          20,
	}

	schedule, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if schedule.ID == 0 {
		t.Fatal("expected schedule ID to be assigned")
	}

	if schedule.DoctorID != req.DoctorID {
		t.Fatalf("expected doctor ID %d, got %d", req.DoctorID, schedule.DoctorID)
	}

	if !schedule.IsActive {
		t.Fatal("expected schedule to be active")
	}
}

func TestCreateScheduleValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateDoctorScheduleRequest{
		{},
		{
			DoctorID:             0,
			Day:                  Monday,
			StartTime:            "08:00",
			EndTime:              "17:00",
			ConsultationDuration: 30,
			MaxPatients:          20,
		},
		{
			DoctorID:             1,
			Day:                  "",
			StartTime:            "08:00",
			EndTime:              "17:00",
			ConsultationDuration: 30,
			MaxPatients:          20,
		},
		{
			DoctorID:             1,
			Day:                  Monday,
			StartTime:            "",
			EndTime:              "17:00",
			ConsultationDuration: 30,
			MaxPatients:          20,
		},
		{
			DoctorID:             1,
			Day:                  Monday,
			StartTime:            "08:00",
			EndTime:              "",
			ConsultationDuration: 30,
			MaxPatients:          20,
		},
		{
			DoctorID:             1,
			Day:                  Monday,
			StartTime:            "08:00",
			EndTime:              "17:00",
			ConsultationDuration: 0,
			MaxPatients:          20,
		},
		{
			DoctorID:             1,
			Day:                  Monday,
			StartTime:            "08:00",
			EndTime:              "17:00",
			ConsultationDuration: 30,
			MaxPatients:          0,
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidSchedule {
			t.Fatalf("expected %v, got %v", ErrInvalidSchedule, err)
		}
	}
}

func TestGetSchedule(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorScheduleRequest{
		DoctorID:             1,
		Day:                  Tuesday,
		StartTime:            "09:00",
		EndTime:              "16:00",
		ConsultationDuration: 20,
		MaxPatients:          18,
	})

	schedule, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if schedule.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, schedule.ID)
	}
}

func TestGetAllSchedules(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(CreateDoctorScheduleRequest{
			DoctorID:             uint(i + 1),
			Day:                  Monday,
			StartTime:            "08:00",
			EndTime:              "17:00",
			ConsultationDuration: 30,
			MaxPatients:          20,
		})
	}

	schedules, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schedules) != 3 {
		t.Fatalf("expected 3 schedules, got %d", len(schedules))
	}
}

func TestGetSchedulesByDoctor(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDoctorScheduleRequest{
		DoctorID:             10,
		Day:                  Wednesday,
		StartTime:            "08:00",
		EndTime:              "17:00",
		ConsultationDuration: 30,
		MaxPatients:          20,
	})

	schedules, err := service.GetByDoctorID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schedules) != 1 {
		t.Fatalf("expected 1 schedule, got %d", len(schedules))
	}
}

func TestGetSchedulesByDay(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDoctorScheduleRequest{
		DoctorID:             1,
		Day:                  Friday,
		StartTime:            "08:00",
		EndTime:              "17:00",
		ConsultationDuration: 30,
		MaxPatients:          20,
	})

	schedules, err := service.GetByDay(Friday)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(schedules) != 1 {
		t.Fatalf("expected 1 schedule, got %d", len(schedules))
	}
}

func TestUpdateSchedule(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorScheduleRequest{
		DoctorID:             1,
		Day:                  Monday,
		StartTime:            "08:00",
		EndTime:              "17:00",
		ConsultationDuration: 30,
		MaxPatients:          20,
	})

	active := false

	updated, err := service.Update(created.ID, UpdateDoctorScheduleRequest{
		StartTime:            "09:00",
		EndTime:              "18:00",
		ConsultationDuration: 45,
		MaxPatients:          15,
		IsActive:             &active,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.StartTime != "09:00" {
		t.Fatal("start time was not updated")
	}

	if updated.EndTime != "18:00" {
		t.Fatal("end time was not updated")
	}

	if updated.ConsultationDuration != 45 {
		t.Fatal("consultation duration was not updated")
	}

	if updated.MaxPatients != 15 {
		t.Fatal("max patients was not updated")
	}

	if updated.IsActive {
		t.Fatal("expected schedule to be inactive")
	}
}

func TestDeleteSchedule(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorScheduleRequest{
		DoctorID:             1,
		Day:                  Sunday,
		StartTime:            "08:00",
		EndTime:              "17:00",
		ConsultationDuration: 30,
		MaxPatients:          20,
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrScheduleNotFound {
		t.Fatalf("expected %v, got %v", ErrScheduleNotFound, err)
	}
}
