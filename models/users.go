package models

type Users struct {
	*Model
	Username string `gorm:"unique"`
	Password string
	Email    string
	Nickname string
	Avatar   string
}
