package domain

import "github.com/gomessguii/todo-list-go/user/models"

type IUserRepository interface {
	GetByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}