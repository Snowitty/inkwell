package models

import "gorm.io/gorm"

type Directory struct {
	gorm.Model
	Name   string
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
