package comments

import "gorm.io/gorm"

type Comment struct {
	gorm.Model

	PostID uint `gorm:"not null;index" json:"post_id"`

	AuthorID uint `gorm:"not null;index" json:"author_id"`

	Content string `gorm:"type:text;not null" json:"content"`

	IsEdited bool `gorm:"default:false" json:"is_edited"`
}
