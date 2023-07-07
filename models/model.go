package models

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy uint
	UpdatedBy uint
	IsDeleted bool
}
