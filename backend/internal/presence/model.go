package presence

import (
	"time"

	"gorm.io/gorm"
)

// Presence represents a user's online presence.
type Presence struct {
	gorm.Model

	UserID uint `gorm:"uniqueIndex;not null" json:"user_id"`

	// Whether the user is currently online.
	IsOnline bool `gorm:"default:false" json:"is_online"`

	// Last time the user was seen online.
	LastSeen time.Time `json:"last_seen"`

	// Optional connection identifier for WebSocket sessions.
	ConnectionID string `gorm:"size:255" json:"connection_id,omitempty"`
}
