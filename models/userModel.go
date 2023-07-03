package models

type User struct {
	ID       uint `gorm:"primarykey"`
	Username string
	Password string
	Email    string
}
