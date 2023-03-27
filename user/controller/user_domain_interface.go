package controller

type IUserDomain interface {
	CreateUser(username, password string) error
	Login(username, password string) error
	IsLoggedIn() bool
	Logout()
}