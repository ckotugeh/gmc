package posts

import "time"

type Post struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	AuthorID    uint      `json:"author_id"`
	CommunityID uint      `json:"community_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	ImageURL    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
}
