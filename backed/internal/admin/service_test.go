package admin

import "testing"

func TestCreateAdmin(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	req := CreateAdminRequest{
		ResourceID:  1,
		Resource:    ResourceUser,
		Action:      ActionCreate,
		Description: "Created a new user",
	}

	admin, err := service.Create(req, 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if admin.ID == 0 {
		t.Fatal("expected admin action ID to be assigned")
	}

	if admin.AdminID != 1 {
		t.Fatalf("expected admin ID 1, got %d", admin.AdminID)
	}

	if admin.Resource != ResourceUser {
		t.Fatalf("expected resource %s, got %s", ResourceUser, admin.Resource)
	}
}

func TestCreateAdminValidation(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	tests := []struct {
		name string
		req  CreateAdminRequest
		id   uint
	}{
		{
			name: "invalid admin id",
			req: CreateAdminRequest{
				ResourceID: 1,
				Resource:   ResourceUser,
				Action:     ActionCreate,
			},
			id: 0,
		},
		{
			name: "missing resource id",
			req: CreateAdminRequest{
				Resource: ResourceUser,
				Action:   ActionCreate,
			},
			id: 1,
		},
		{
			name: "missing resource",
			req: CreateAdminRequest{
				ResourceID: 1,
				Action:     ActionCreate,
			},
			id: 1,
		},
		{
			name: "missing action",
			req: CreateAdminRequest{
				ResourceID: 1,
				Resource:   ResourceUser,
			},
			id: 1,
		},
	}

	for _, tt := range tests {
		_, err := service.Create(tt.req, tt.id)
		if err != ErrInvalidAdmin {
			t.Fatalf("%s: expected %v, got %v", tt.name, ErrInvalidAdmin, err)
		}
	}
}

func TestGetAdmin(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	admin, _ := service.Create(CreateAdminRequest{
		ResourceID: 1,
		Resource:   ResourceUser,
		Action:     ActionCreate,
	}, 1)

	found, err := service.GetByID(admin.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if found.ID != admin.ID {
		t.Fatalf("expected ID %d, got %d", admin.ID, found.ID)
	}
}

func TestGetAllAdmins(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	for i := 0; i < 3; i++ {
		_, _ = service.Create(CreateAdminRequest{
			ResourceID: uint(i + 1),
			Resource:   ResourceUser,
			Action:     ActionCreate,
		}, 1)
	}

	admins, err := service.GetAll()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(admins) != 3 {
		t.Fatalf("expected 3 admin actions, got %d", len(admins))
	}
}

func TestGetByAdminID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.Create(CreateAdminRequest{
		ResourceID: 1,
		Resource:   ResourceUser,
		Action:     ActionCreate,
	}, 1)

	_, _ = service.Create(CreateAdminRequest{
		ResourceID: 2,
		Resource:   ResourceHospital,
		Action:     ActionApprove,
	}, 1)

	admins, err := service.GetByAdminID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(admins) != 2 {
		t.Fatalf("expected 2 admin actions, got %d", len(admins))
	}
}

func TestUpdateAdmin(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	admin, _ := service.Create(CreateAdminRequest{
		ResourceID: 1,
		Resource:   ResourceUser,
		Action:     ActionCreate,
	}, 1)

	updated, err := service.Update(admin.ID, UpdateAdminRequest{
		Action:      ActionSuspend,
		Description: "Suspended user account",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Action != ActionSuspend {
		t.Fatalf("expected action %s, got %s", ActionSuspend, updated.Action)
	}

	if updated.Description != "Suspended user account" {
		t.Fatal("description not updated")
	}
}

func TestDeleteAdmin(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	admin, _ := service.Create(CreateAdminRequest{
		ResourceID: 1,
		Resource:   ResourceUser,
		Action:     ActionDelete,
	}, 1)

	if err := service.Delete(admin.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := service.GetByID(admin.ID)
	if err != ErrAdminNotFound {
		t.Fatalf("expected %v, got %v", ErrAdminNotFound, err)
	}
}

func TestDashboardStats(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	stats, err := service.GetDashboardStats()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if stats.TotalUsers == 0 {
		t.Fatal("expected dashboard statistics")
	}
}
