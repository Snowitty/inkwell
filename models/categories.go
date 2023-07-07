package models

type Categories struct {
	*Model
	Name        string
	Description string
	CoverImgURL string
	Articles    []*Articles `gorm:"many2many:article_categories;"`
}
