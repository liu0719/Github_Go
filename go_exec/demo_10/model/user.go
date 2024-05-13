package model

type User struct {
	Name    string
	Age     uint
	Tele    string
	Address string
}
type Team struct {
	Name    string
	User    *User
	Address string
}
