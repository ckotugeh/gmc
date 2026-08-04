package users

import "testing"

func TestCreateUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
		Phone:     "0712345678",
		Role:      "doctor",
	}

	user, err := service.CreateUser(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID == 0 {
		t.Fatal("expected user ID to be set")
	}

	if user.Email != req.Email {
		t.Fatalf("expected email %s, got %s", req.Email, user.Email)
	}

	if user.Password == req.Password {
		t.Fatal("password should be hashed")
	}
}

func TestCreateUserDuplicateEmail(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	}

	_, _ = service.CreateUser(req)

	_, err := service.CreateUser(req)
	if err == nil {
		t.Fatal("expected duplicate email error")
	}
}

func TestGetUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateUser(CreateUserRequest{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "password123",
	})

	user, err := service.GetUser(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.ID != created.ID {
		t.Fatal("unexpected user returned")
	}
}

func TestGetAllUsers(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreateUser(CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	})

	_, _ = service.CreateUser(CreateUserRequest{
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "jane@example.com",
		Password:  "password123",
	})

	users, err := service.GetAllUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

func TestGetDoctors(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreateUser(CreateUserRequest{
		FirstName: "Doctor",
		LastName:  "One",
		Email:     "doctor@example.com",
		Password:  "password123",
		Role:      "doctor",
	})

	_, _ = service.CreateUser(CreateUserRequest{
		FirstName: "Patient",
		LastName:  "One",
		Email:     "patient@example.com",
		Password:  "password123",
		Role:      "patient",
	})

	doctors, err := service.GetDoctors()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(doctors) != 1 {
		t.Fatalf("expected 1 doctor, got %d", len(doctors))
	}
}

func TestGetPatients(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreateUser(CreateUserRequest{
		FirstName: "Doctor",
		LastName:  "One",
		Email:     "doctor@example.com",
		Password:  "password123",
		Role:      "doctor",
	})

	_, _ = service.CreateUser(CreateUserRequest{
		FirstName: "Patient",
		LastName:  "One",
		Email:     "patient@example.com",
		Password:  "password123",
		Role:      "patient",
	})

	patients, err := service.GetPatients()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(patients) != 1 {
		t.Fatalf("expected 1 patient, got %d", len(patients))
	}
}

func TestUpdateUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateUser(CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	})

	firstName := "Johnny"
	phone := "0700000000"

	user, err := service.UpdateUser(created.ID, UpdateUserRequest{
		FirstName: &firstName,
		Phone:     &phone,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.FirstName != firstName {
		t.Fatal("first name not updated")
	}

	if user.Phone != phone {
		t.Fatal("phone not updated")
	}
}

func TestDeleteUser(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreateUser(CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Password:  "password123",
	})

	if err := service.DeleteUser(created.ID); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := service.GetUser(created.ID); err == nil {
		t.Fatal("expected user to be deleted")
	}
}
