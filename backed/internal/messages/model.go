package messages

import "gorm.io/gorm"

const (
	MessageTypeText  = "text"
	MessageTypeImage = "image"
	MessageTypeFile  = "file"
)

type Message struct {
	gorm.Model

	// Conversation participants
	SenderID   uint `json:"sender_id" gorm:"not null;index"`
	ReceiverID uint `json:"receiver_id" gorm:"not null;index"`

	// Message content
	Content     string `json:"content" gorm:"type:text;not null"`
	MessageType string `json:"message_type" gorm:"type:varchar(20);default:'text'"`

	// Optional attachment (image, pdf, etc.)
	AttachmentURL string `json:"attachment_url"`

	// Reply to another message (optional)
	ReplyToID *uint `json:"reply_to_id"`

	// Status
	IsRead   bool `json:"is_read" gorm:"default:false"`
	IsEdited bool `json:"is_edited" gorm:"default:false"`

	// Relationships (enable later if needed)
	// Sender   auth.User `gorm:"foreignKey:SenderID"`
	// Receiver auth.User `gorm:"foreignKey:ReceiverID"`
	// ReplyTo  *Message  `gorm:"foreignKey:ReplyToID"`
}
