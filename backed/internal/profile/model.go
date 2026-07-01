package profile

import "gorm.io/gorm"

type Profile struct {
	gorm.Model

	UserID uint `gorm:"uniqueIndex;not null"`

	FullName string `gorm:"size:100;not null"`

	Specialization string `gorm:"size:100;not null"`
	Hospital       string `gorm:"size:150"`

	Country string `gorm:"size:100"`
	City    string `gorm:"size:100"`

	YearsExperience int

	LicenseNumber string `gorm:"size:100"`
	Education     string `gorm:"size:255"`

	Languages string `gorm:"size:255"`

	Bio string `gorm:"type:text"`

	ProfileImageURL string `gorm:"size:255"`
}
