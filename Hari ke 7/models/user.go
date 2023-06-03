package models

type User struct {
	Id int `gorm:"primaryKey autoIncrement" json:"id"`
	Name string `json:"name"`
	PhotoProfile string `json:"photoProfile"`
	Password string `json:"password"`
}