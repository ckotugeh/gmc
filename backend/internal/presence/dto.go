package presence

import "time"

// UpdatePresenceRequest represents a request to update a user's presence.
type UpdatePresenceRequest struct {
	IsOnline bool `json:"is_online"`
}

// PresenceResponse represents a user's presence information.
type PresenceResponse struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	IsOnline     bool      `json:"is_online"`
	LastSeen     time.Time `json:"last_seen"`
	ConnectionID string    `json:"connection_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
