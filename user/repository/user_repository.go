package repository

import (
	"fmt"

	"github.com/gomessguii/todo-list-go/user/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	userTransfer := user.ToTransfer()
	foundUser, _ := r.GetByUsername(userTransfer.Username)
	if foundUser.IsValid() {
		return fmt.Errorf("username already registred")
	}
	return r.db.Create(userTransfer).Error
}

func (r *userRepository) GetByUsername(username string) (response *models.User, err error) {
	err = r.db.Find(&response, "username LIKE", username).Error
	return
}