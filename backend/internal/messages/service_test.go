package messages

import (
	"testing"
)

func setupService() Service {
	repo := NewMockRepository()
	return NewService(repo)
}

func createTestMessage(service Service, senderID, receiverID uint) *Message {
	msg, _ := service.CreateMessage(senderID, &CreateMessageRequest{
		ReceiverID: receiverID,
		Content:    "Hello",
	})
	return msg
}

func TestCreateMessage(t *testing.T) {
	s := setupService()

	msg, err := s.CreateMessage(1, &CreateMessageRequest{
		ReceiverID: 2,
		Content:    "Hello Doctor",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if msg.ID == 0 {
		t.Error("expected message ID to be set")
	}

	if msg.Content != "Hello Doctor" {
		t.Error("message content mismatch")
	}
}

func TestCreateMessageValidation(t *testing.T) {
	s := setupService()

	_, err := s.CreateMessage(1, &CreateMessageRequest{
		ReceiverID: 0,
		Content:    "",
	})

	if err == nil {
		t.Error("expected validation error")
	}
}

func TestSelfMessaging(t *testing.T) {
	s := setupService()

	_, err := s.CreateMessage(1, &CreateMessageRequest{
		ReceiverID: 1,
		Content:    "Hello myself",
	})

	if err == nil {
		t.Error("expected error when sending message to self")
	}
}

func TestUpdateMessage(t *testing.T) {
	s := setupService()

	msg := createTestMessage(s, 1, 2)

	updated, err := s.UpdateMessage(msg.ID, 1, &UpdateMessageRequest{
		Content: "Updated message",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !updated.IsEdited {
		t.Error("expected message to be marked as edited")
	}

	if updated.Content != "Updated message" {
		t.Error("content not updated")
	}
}

func TestUpdateMessageUnauthorized(t *testing.T) {
	s := setupService()

	msg := createTestMessage(s, 1, 2)

	_, err := s.UpdateMessage(msg.ID, 99, &UpdateMessageRequest{
		Content: "Hack attempt",
	})

	if err == nil {
		t.Error("expected unauthorized error")
	}
}

func TestDeleteMessage(t *testing.T) {
	s := setupService()

	msg := createTestMessage(s, 1, 2)

	err := s.DeleteMessage(msg.ID, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err = s.GetMessage(msg.ID)
	if err == nil {
		t.Error("expected message to be deleted")
	}
}

func TestDeleteMessageUnauthorized(t *testing.T) {
	s := setupService()

	msg := createTestMessage(s, 1, 2)

	err := s.DeleteMessage(msg.ID, 99)
	if err == nil {
		t.Error("expected unauthorized delete error")
	}
}

func TestMarkAsRead(t *testing.T) {
	s := setupService()

	msg := createTestMessage(s, 1, 2)

	err := s.MarkAsRead(msg.ID, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updated, _ := s.GetMessage(msg.ID)
	if !updated.IsRead {
		t.Error("expected message to be marked as read")
	}
}

func TestMarkAsReadUnauthorized(t *testing.T) {
	s := setupService()

	msg := createTestMessage(s, 1, 2)

	err := s.MarkAsRead(msg.ID, 1)
	if err == nil {
		t.Error("expected unauthorized error")
	}
}

func TestGetConversation(t *testing.T) {
	s := setupService()

	_, _ = s.CreateMessage(1, &CreateMessageRequest{
		ReceiverID: 2,
		Content:    "Hi",
	})

	_, _ = s.CreateMessage(2, &CreateMessageRequest{
		ReceiverID: 1,
		Content:    "Hello back",
	})

	convo, err := s.GetConversation(1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(convo) != 2 {
		t.Errorf("expected 2 messages, got %d", len(convo))
	}
}
