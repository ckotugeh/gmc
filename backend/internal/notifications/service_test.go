package notifications

import "testing"

func setupService() Service {
	repo := NewMockRepository()
	return NewService(repo)
}

func createTestNotification(s Service, userID uint) *Notification {
	n, _ := s.CreateNotification(&CreateNotificationRequest{
		UserID:  userID,
		Type:    NotificationTypeMessage,
		Title:   "New Message",
		Message: "You have received a new message.",
	})

	return n
}

func TestCreateNotification(t *testing.T) {
	s := setupService()

	n, err := s.CreateNotification(&CreateNotificationRequest{
		UserID:  1,
		Type:    NotificationTypeComment,
		Title:   "New Comment",
		Message: "Someone commented on your post.",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if n.ID == 0 {
		t.Error("expected notification ID to be assigned")
	}

	if n.IsRead {
		t.Error("new notification should be unread")
	}
}

func TestCreateNotificationValidation(t *testing.T) {
	s := setupService()

	_, err := s.CreateNotification(&CreateNotificationRequest{})

	if err == nil {
		t.Error("expected validation error")
	}
}

func TestGetNotification(t *testing.T) {
	s := setupService()

	n := createTestNotification(s, 1)

	result, err := s.GetNotification(n.ID, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.ID != n.ID {
		t.Error("notification ID mismatch")
	}
}

func TestGetNotificationUnauthorized(t *testing.T) {
	s := setupService()

	n := createTestNotification(s, 1)

	_, err := s.GetNotification(n.ID, 2)

	if err == nil {
		t.Error("expected unauthorized error")
	}
}

func TestGetUserNotifications(t *testing.T) {
	s := setupService()

	createTestNotification(s, 1)
	createTestNotification(s, 1)
	createTestNotification(s, 2)

	list, err := s.GetUserNotifications(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 2 {
		t.Errorf("expected 2 notifications, got %d", len(list))
	}
}

func TestMarkAsRead(t *testing.T) {
	s := setupService()

	n := createTestNotification(s, 1)

	err := s.MarkAsRead(n.ID, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updated, _ := s.GetNotification(n.ID, 1)

	if !updated.IsRead {
		t.Error("expected notification to be marked as read")
	}
}

func TestMarkAsReadUnauthorized(t *testing.T) {
	s := setupService()

	n := createTestNotification(s, 1)

	err := s.MarkAsRead(n.ID, 2)

	if err == nil {
		t.Error("expected unauthorized error")
	}
}

func TestMarkAllAsRead(t *testing.T) {
	s := setupService()

	createTestNotification(s, 1)
	createTestNotification(s, 1)

	err := s.MarkAllAsRead(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	unread, err := s.GetUnreadNotifications(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(unread) != 0 {
		t.Errorf("expected 0 unread notifications, got %d", len(unread))
	}
}

func TestDeleteNotification(t *testing.T) {
	s := setupService()

	n := createTestNotification(s, 1)

	err := s.DeleteNotification(n.ID, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = s.GetNotification(n.ID, 1)
	if err == nil {
		t.Error("expected notification to be deleted")
	}
}

func TestDeleteNotificationUnauthorized(t *testing.T) {
	s := setupService()

	n := createTestNotification(s, 1)

	err := s.DeleteNotification(n.ID, 2)

	if err == nil {
		t.Error("expected unauthorized error")
	}
}
