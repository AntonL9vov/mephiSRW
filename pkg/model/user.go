package model

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Email    string `json:"email" binding:"required"`
	phone    string `json:"group_id"`
	Password string `json:"password" binding:"required"`
}
