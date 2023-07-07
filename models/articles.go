package models

type Articles struct {
	*Model
	Title       string
	Description string
	CoverImgURL string
	Content     string
	Tags        []*Tags       `gorm:"many2many:article_tags;"`
	Categories  []*Categories `gorm:"many2many:article_categories;"`
}
