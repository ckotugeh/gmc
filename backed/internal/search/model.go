package search

import (
	"time"

	"gorm.io/gorm"
)

// SearchType represents the category of a search result.
type SearchType string

const (
	TypeDoctor      SearchType = "doctor"
	TypePatient     SearchType = "patient"
	TypeHospital    SearchType = "hospital"
	TypeCommunity   SearchType = "community"
	TypePost        SearchType = "post"
	TypeAppointment SearchType = "appointment"
)

// Search represents a user's search history.
type Search struct {
	ID uint `gorm:"primaryKey" json:"id"`

	UserID uint `gorm:"not null;index" json:"user_id"`

	Query string `gorm:"size:255;not null;index" json:"query"`

	Type SearchType `gorm:"type:varchar(30);not null" json:"type"`

	ResultCount int `gorm:"default:0" json:"result_count"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
