package reactions

import "gorm.io/gorm"

const (
	ReactionLike    = "like"
	ReactionDislike = "dislike"
)

type Reaction struct {
	gorm.Model

	PostID uint `json:"post_id" gorm:"not null;index"`
	UserID uint `json:"user_id" gorm:"not null;index"`

	// Current reaction: "like" or "dislike"
	ReactionType string `json:"reaction_type" gorm:"type:varchar(20);not null"`

	// Relationships
	// Post posts.Post `gorm:"foreignKey:PostID"`
	// User auth.User `gorm:"foreignKey:UserID"`
}
