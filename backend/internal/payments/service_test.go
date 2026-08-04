package payments

import "testing"

func TestCreatePayment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     1,
		DoctorID:      2,
		Amount:        5000,
		Currency:      CurrencyKES,
		Method:        MethodMobileMoney,
		Description:   "Consultation payment",
	}

	payment, err := service.Create(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if payment.ID == 0 {
		t.Fatal("expected payment ID to be assigned")
	}

	if payment.Status != StatusPending {
		t.Fatalf("expected status %s, got %s", StatusPending, payment.Status)
	}
}

func TestCreatePaymentValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []CreatePaymentRequest{
		{
			AppointmentID: 0,
			PatientID:     1,
			DoctorID:      1,
			Amount:        100,
			Method:        MethodCash,
		},
		{
			AppointmentID: 1,
			PatientID:     0,
			DoctorID:      1,
			Amount:        100,
			Method:        MethodCash,
		},
		{
			AppointmentID: 1,
			PatientID:     1,
			DoctorID:      0,
			Amount:        100,
			Method:        MethodCash,
		},
		{
			AppointmentID: 1,
			PatientID:     1,
			DoctorID:      1,
			Amount:        0,
			Method:        MethodCash,
		},
	}

	for _, req := range tests {
		_, err := service.Create(req)
		if err != ErrInvalidPayment {
			t.Fatalf("expected %v, got %v", ErrInvalidPayment, err)
		}
	}
}

func TestGetPayment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	payment, _ := service.Create(CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     1,
		DoctorID:      2,
		Amount:        1000,
		Method:        MethodCash,
	})

	found, err := service.GetByID(payment.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != payment.ID {
		t.Fatalf("expected payment ID %d, got %d", payment.ID, found.ID)
	}
}

func TestGetAllPayments(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(CreatePaymentRequest{
			AppointmentID: uint(i + 1),
			PatientID:     1,
			DoctorID:      2,
			Amount:        100,
			Method:        MethodCash,
		})
	}

	payments, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(payments) != 3 {
		t.Fatalf("expected 3 payments, got %d", len(payments))
	}
}

func TestGetByPatientID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     10,
		DoctorID:      2,
		Amount:        100,
		Method:        MethodCash,
	})

	payments, err := service.GetByPatientID(10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(payments) != 1 {
		t.Fatalf("expected 1 payment, got %d", len(payments))
	}
}

func TestGetByDoctorID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     1,
		DoctorID:      20,
		Amount:        100,
		Method:        MethodCash,
	})

	payments, err := service.GetByDoctorID(20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(payments) != 1 {
		t.Fatalf("expected 1 payment, got %d", len(payments))
	}
}

func TestUpdatePayment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	payment, _ := service.Create(CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     1,
		DoctorID:      2,
		Amount:        1000,
		Method:        MethodCash,
	})

	updated, err := service.Update(payment.ID, UpdatePaymentRequest{
		Status:               StatusPaid,
		TransactionReference: "TXN123456",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Status != StatusPaid {
		t.Fatalf("expected status %s, got %s", StatusPaid, updated.Status)
	}

	if updated.TransactionReference != "TXN123456" {
		t.Fatal("transaction reference not updated")
	}

	if updated.PaidAt == nil {
		t.Fatal("expected PaidAt to be set")
	}
}

func TestDeletePayment(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	payment, _ := service.Create(CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     1,
		DoctorID:      2,
		Amount:        100,
		Method:        MethodCash,
	})

	if err := service.Delete(payment.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(payment.ID)
	if err != ErrPaymentNotFound {
		t.Fatalf("expected %v, got %v", ErrPaymentNotFound, err)
	}
}

func TestPaymentSummary(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	payment, _ := service.Create(CreatePaymentRequest{
		AppointmentID: 1,
		PatientID:     1,
		DoctorID:      2,
		Amount:        500,
		Method:        MethodCash,
	})

	_, _ = service.Update(payment.ID, UpdatePaymentRequest{
		Status: StatusPaid,
	})

	summary, err := service.GetSummary()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if summary.TotalPayments != 1 {
		t.Fatalf("expected 1 payment, got %d", summary.TotalPayments)
	}

	if summary.TotalRevenue != 500 {
		t.Fatalf("expected revenue 500, got %f", summary.TotalRevenue)
	}
}
