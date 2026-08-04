package websockets

import "time"

// MessageType represents the type of websocket message.
type MessageType string

const (
	// Chat messages
	MessageTypeChat MessageType = "chat"

	// Notifications
	MessageTypeNotification MessageType = "notification"

	// User presence
	MessageTypePresence MessageType = "presence"

	// Typing indicators
	MessageTypeTyping MessageType = "typing"

	// Read receipts
	MessageTypeReadReceipt MessageType = "read_receipt"

	// Ping/Pong (heartbeat)
	MessageTypePing MessageType = "ping"
	MessageTypePong MessageType = "pong"

	// Error messages
	MessageTypeError MessageType = "error"
)

// Message is the standard websocket payload exchanged between
// the client and server.
type Message struct {
	Type MessageType `json:"type"`

	// Sender information
	SenderID uint `json:"sender_id,omitempty"`

	// Recipient information
	ReceiverID uint `json:"receiver_id,omitempty"`

	// Optional notification recipient
	UserID uint `json:"user_id,omitempty"`

	// Content
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`

	// Reference to another resource
	ReferenceType string `json:"reference_type,omitempty"`
	ReferenceID   uint   `json:"reference_id,omitempty"`

	// Conversation helpers
	ConversationID uint `json:"conversation_id,omitempty"`

	// Typing indicator
	IsTyping bool `json:"is_typing,omitempty"`

	// Presence
	IsOnline bool `json:"is_online,omitempty"`

	// Read receipt
	IsRead bool `json:"is_read,omitempty"`

	// Error information
	Error string `json:"error,omitempty"`

	// Timestamp
	Timestamp time.Time `json:"timestamp"`
}
