package models

type User struct {
	Id string
	Login string
	Password string
}

func createUser() *User {
	return &User{}
}
