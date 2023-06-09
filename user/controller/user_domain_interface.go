package controller

import "github.com/gomessguii/todo-list-go/user/models"

type IUserDomain interface {
	CreateUser(user *models.User) error
	Login(username, password string) error
	IsLoggedIn() bool
	Logout()
}