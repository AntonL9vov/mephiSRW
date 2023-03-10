package model

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}
