package doctor_reviews

import "testing"

func TestCreateDoctorReview(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 100,
		Rating:        5,
		Title:         "Excellent Doctor",
		Comment:       "Very professional and caring.",
	}

	review, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if review.ID == 0 {
		t.Fatal("expected review ID to be assigned")
	}

	if review.Rating != 5 {
		t.Fatalf("expected rating 5, got %d", review.Rating)
	}

	if !review.IsPublished {
		t.Fatal("expected review to be published by default")
	}

	if review.IsAnonymous {
		t.Fatal("expected review to be non-anonymous by default")
	}
}

func TestCreateDoctorReviewValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreateDoctorReviewRequest{
		{},
		{
			DoctorID:      1,
			PatientID:     1,
			AppointmentID: 1,
			Rating:        0,
			Comment:       "Test",
		},
		{
			DoctorID:      1,
			PatientID:     1,
			AppointmentID: 1,
			Rating:        6,
			Comment:       "Test",
		},
		{
			DoctorID:      1,
			PatientID:     1,
			AppointmentID: 1,
			Rating:        5,
			Comment:       "",
		},
	}

	for _, tc := range tests {
		_, err := service.Create(tc)
		if err != ErrInvalidDoctorReview {
			t.Fatalf("expected %v, got %v", ErrInvalidDoctorReview, err)
		}
	}
}

func TestDuplicateAppointmentReview(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     1,
		AppointmentID: 10,
		Rating:        5,
		Comment:       "Excellent",
	}

	_, _ = service.Create(req)

	_, err := service.Create(req)
	if err != ErrDuplicateAppointmentReview {
		t.Fatalf("expected %v, got %v", ErrDuplicateAppointmentReview, err)
	}
}

func TestGetDoctorReview(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 101,
		Rating:        4,
		Comment:       "Very good",
	})

	review, err := service.GetByID(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if review.ID != created.ID {
		t.Fatalf("expected ID %d, got %d", created.ID, review.ID)
	}
}

func TestGetByAppointmentID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 200,
		Rating:        5,
		Comment:       "Excellent",
	})

	review, err := service.GetByAppointmentID(created.AppointmentID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if review.AppointmentID != created.AppointmentID {
		t.Fatal("appointment ID mismatch")
	}
}

func TestGetByDoctorID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     1,
		AppointmentID: 1,
		Rating:        5,
		Comment:       "Excellent",
	})

	_, _ = service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 2,
		Rating:        4,
		Comment:       "Very good",
	})

	reviews, err := service.GetByDoctorID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(reviews) != 2 {
		t.Fatalf("expected 2 reviews, got %d", len(reviews))
	}
}

func TestGetPublishedReviews(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     1,
		AppointmentID: 1,
		Rating:        5,
		Comment:       "Excellent",
	})

	published := false

	_, _ = service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     2,
		AppointmentID: 2,
		Rating:        4,
		Comment:       "Hidden review",
		IsPublished:   &published,
	})

	reviews, err := service.GetPublishedByDoctorID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(reviews) != 1 {
		t.Fatalf("expected 1 published review, got %d", len(reviews))
	}
}

func TestUpdateDoctorReview(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     1,
		AppointmentID: 1,
		Rating:        5,
		Comment:       "Excellent",
	})

	anonymous := true

	updated, err := service.Update(created.ID, UpdateDoctorReviewRequest{
		Rating:      4,
		Title:       "Updated",
		Comment:     "Still very good",
		IsAnonymous: &anonymous,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Rating != 4 {
		t.Fatal("rating was not updated")
	}

	if updated.Title != "Updated" {
		t.Fatal("title was not updated")
	}

	if !updated.IsAnonymous {
		t.Fatal("expected anonymous review")
	}
}

func TestDeleteDoctorReview(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.Create(CreateDoctorReviewRequest{
		DoctorID:      1,
		PatientID:     1,
		AppointmentID: 10,
		Rating:        5,
		Comment:       "Excellent",
	})

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(created.ID)
	if err != ErrDoctorReviewNotFound {
		t.Fatalf("expected %v, got %v", ErrDoctorReviewNotFound, err)
	}
}
