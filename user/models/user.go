package models

type User struct {
	id int
	username string
	password string
}

func (u *User) CheckPassword(password string) bool {
	return u.password == password
}