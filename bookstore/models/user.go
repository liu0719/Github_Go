package models

type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

func (User) TableName() string {
	return "users"
}
