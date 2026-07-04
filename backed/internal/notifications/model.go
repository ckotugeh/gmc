package notifications

import "gorm.io/gorm"

const (
	NotificationTypeMessage    = "message"
	NotificationTypeComment    = "comment"
	NotificationTypeReaction   = "reaction"
	NotificationTypeCommunity  = "community"
	NotificationTypePost       = "post"
	NotificationTypeSystem     = "system"
	NotificationTypeMention    = "mention"
	NotificationTypeFollow     = "follow"
	NotificationTypeInvitation = "invitation"
)

type Notification struct {
	gorm.Model

	// User receiving the notification
	UserID uint `json:"user_id" gorm:"not null;index"`

	// User who triggered the notification (optional for system notifications)
	ActorID *uint `json:"actor_id,omitempty" gorm:"index"`

	// Notification type
	Type string `json:"type" gorm:"type:varchar(50);not null;index"`

	// Notification content
	Title   string `json:"title" gorm:"size:255;not null"`
	Message string `json:"message" gorm:"type:text;not null"`

	// Related resource
	ReferenceID   *uint  `json:"reference_id,omitempty"`
	ReferenceType string `json:"reference_type,omitempty" gorm:"size:50"`

	// Read status
	IsRead bool `json:"is_read" gorm:"default:false"`

	// Optional deep link for frontend navigation
	ActionURL string `json:"action_url,omitempty"`

	// Future relationships (enable later if needed)
	// User  auth.User `gorm:"foreignKey:UserID"`
	// Actor auth.User `gorm:"foreignKey:ActorID"`
}
