package models

import (
	"gorm.io/gorm"
)

type ShortLink struct {
	gorm.Model
	// ID          uint
	originalUrl string
	shortLink   string
	// CreatedAt   time.Time
	// UpdatedAt   time.Time
}
