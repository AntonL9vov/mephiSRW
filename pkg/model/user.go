package model

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
