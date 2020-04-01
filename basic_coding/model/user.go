package model

type UserInterface interface {
	Insert() User
}

type User struct {
	Id      int
	Name    string
	Address string
}

func (u *User) Insert() User {
	return *u
}
