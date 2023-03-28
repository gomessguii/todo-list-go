package domain

import (
	"fmt"

	"github.com/gomessguii/todo-list-go/user/models"
)

type IUserDomain interface {
	CreateUser(user *models.User) error
	Login(username, password string) error
	IsLoggedIn() bool
	Logout()
}

type userDomain struct {
	loggedUser *models.User
	userRepository IUserRepository
}

func New(userRepository IUserRepository) IUserDomain {
	return &userDomain{userRepository: userRepository}
}

func (u *userDomain) CreateUser(user *models.User) error {
	if !user.IsValid() {
		return fmt.Errorf("user is not valid")
	}
	return u.userRepository.CreateUser(user)
}

func (u *userDomain) IsLoggedIn() bool {
	return u.loggedUser.IsValid()
}

func (u *userDomain) Login(username string, password string) error {
	user, err := u.userRepository.GetByUsername(username)
	if err != nil {
		return err
	}
	if !user.CheckPassword(password) {
		return fmt.Errorf("incorrect password")
	}
	return nil
}

func (u *userDomain) Logout() {
	u.loggedUser = nil
}