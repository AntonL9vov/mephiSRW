package mephiSRW

type User struct {
	Id         int    `json:"-"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	GroupId    int    `json:"group_id"`
	Password   string `json:"password"`
}
