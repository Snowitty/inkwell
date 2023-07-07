package models

type ArticleCategory struct {
	ArticleID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}
