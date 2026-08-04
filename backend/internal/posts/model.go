package posts

import (
	"time"

	"gorm.io/gorm"
)

// Post represents a discussion post within a community.
type Post struct {
	gorm.Model

	// Relationships
	CommunityID uint `gorm:"not null;index"`
	AuthorID    uint `gorm:"not null;index"`

	// Main content
	Title       string `gorm:"size:255;not null"`
	Content     string `gorm:"type:text;not null"`
	ContentType string `gorm:"size:50;default:'markdown'"`

	// Optional cover image
	ImageURL string

	// Visibility
	IsAnonymous bool `gorm:"default:false"`
	IsEdited    bool `gorm:"default:false"`
	IsPinned    bool `gorm:"default:false"`
	IsLocked    bool `gorm:"default:false"`

	// Statistics
	LikesCount    int `gorm:"default:0"`
	CommentsCount int `gorm:"default:0"`
	ViewsCount    int `gorm:"default:0"`

	// Relationships
	Attachments []Attachment `gorm:"foreignKey:PostID"`
	Tags        []Tag        `gorm:"many2many:post_tags;"`
	Polls       []Poll       `gorm:"foreignKey:PostID"`
}

// ------------------------------------------------------
// Attachments
// ------------------------------------------------------

type Attachment struct {
	gorm.Model

	PostID uint `gorm:"not null;index"`

	FileName string
	FileURL  string
	FileType string
	FileSize int64
}

// ------------------------------------------------------
// Tags
// ------------------------------------------------------

type Tag struct {
	gorm.Model

	Name string `gorm:"uniqueIndex;size:100"`
}

// ------------------------------------------------------
// Polls
// ------------------------------------------------------

type Poll struct {
	gorm.Model

	PostID uint `gorm:"not null;index"`

	Question string `gorm:"size:255;not null"`

	ExpiresAt *time.Time

	Options []PollOption `gorm:"foreignKey:PollID"`
}

// ------------------------------------------------------
// Poll Options
// ------------------------------------------------------

type PollOption struct {
	gorm.Model

	PollID uint `gorm:"not null;index"`

	Text  string `gorm:"size:255;not null"`
	Votes int    `gorm:"default:0"`
}
