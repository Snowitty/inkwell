package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string
	Content     string
	TagID       uint
	Tag         Tag `gorm:"foreignKey:TagID"`
	DirectoryID uint
	Directory   Directory `gorm:"foreignKey:DirectoryID"`
	CreatedAt   time.Time
	UpdateAt    time.Time
}
