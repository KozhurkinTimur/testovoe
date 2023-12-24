package models

type User struct {
	Id string
	login string
	password string
}

func createUser() *User {
	return &User{}
}
