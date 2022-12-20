package models

import (
	"gorm.io/gorm"
)

type ShortLink struct {
	gorm.Model
	OriginalUrl string
	ShortCode   string `gorm:"uniqueIndex"`
}
