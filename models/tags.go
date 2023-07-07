package models

type Tags struct {
	*Model
	Name        string
	Description string
	Articles    []*Articles `gorm:"many2many:article_tags;"`
}
