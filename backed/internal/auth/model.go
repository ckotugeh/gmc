package auth

import "time"

type User struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	FullName           string    `json:"full_name"`
	Email              string    `gorm:"unique" json:"email"`
	Password           string    `json:"-"`
	Role               string    `json:"role"`
	VerificationStatus string    `json:"verification_status"`
	CreatedAt          time.Time `json:"created_at"`
}
