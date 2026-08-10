package presence

import "testing"

func TestCreatePresence(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	presence, err := service.CreatePresence(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if presence.UserID != 1 {
		t.Errorf("expected user ID 1, got %d", presence.UserID)
	}

	if presence.IsOnline {
		t.Errorf("expected user to be offline by default")
	}
}

func TestCreatePresenceDuplicate(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreatePresence(1)

	_, err := service.CreatePresence(1)
	if err == nil {
		t.Fatal("expected duplicate presence error")
	}
}

func TestGetPresence(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	created, _ := service.CreatePresence(1)

	presence, err := service.GetPresence(created.ID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if presence.ID != created.ID {
		t.Errorf("expected ID %d, got %d", created.ID, presence.ID)
	}
}

func TestGetPresenceByUserID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreatePresence(15)

	presence, err := service.GetPresenceByUserID(15)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if presence.UserID != 15 {
		t.Errorf("expected user ID 15, got %d", presence.UserID)
	}
}

func TestUpdatePresenceOnline(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreatePresence(1)

	presence, err := service.UpdatePresence(1, UpdatePresenceRequest{
		IsOnline: true,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !presence.IsOnline {
		t.Errorf("expected user to be online")
	}
}

func TestUpdatePresenceOffline(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreatePresence(1)

	_, _ = service.UpdatePresence(1, UpdatePresenceRequest{
		IsOnline: true,
	})

	presence, err := service.UpdatePresence(1, UpdatePresenceRequest{
		IsOnline: false,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if presence.IsOnline {
		t.Errorf("expected user to be offline")
	}

	if presence.LastSeen.IsZero() {
		t.Errorf("expected last_seen to be updated")
	}
}

func TestSetConnectionID(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreatePresence(1)

	presence, err := service.SetConnectionID(1, "conn-123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if presence.ConnectionID != "conn-123" {
		t.Errorf("expected connection ID conn-123, got %s", presence.ConnectionID)
	}
}

func TestGetOnlineUsers(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	_, _ = service.CreatePresence(1)
	_, _ = service.CreatePresence(2)

	_, _ = service.UpdatePresence(1, UpdatePresenceRequest{
		IsOnline: true,
	})

	users, err := service.GetOnlineUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 1 {
		t.Fatalf("expected 1 online user, got %d", len(users))
	}

	if users[0].UserID != 1 {
		t.Errorf("expected user ID 1")
	}
}

func TestDeletePresence(t *testing.T) {
	repo := NewMockRepository()
	service := NewService(repo)

	presence, _ := service.CreatePresence(1)

	if err := service.DeletePresence(presence.ID); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err := service.GetPresence(presence.ID)
	if err == nil {
		t.Fatal("expected presence to be deleted")
	}
}
