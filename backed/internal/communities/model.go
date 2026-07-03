package communities

import "gorm.io/gorm"

type Community struct {
	gorm.Model

	Name        string `gorm:"size:100;uniqueIndex;not null"`
	Description string `gorm:"type:text"`
	Category    string `gorm:"size:100;not null"`

	CreatorID uint `gorm:"not null"`

	BannerURL string `gorm:"size:255"`

	IsPrivate bool `gorm:"default:false"`
}