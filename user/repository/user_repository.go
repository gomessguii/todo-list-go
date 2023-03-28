package repository

import (
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
	return r.db.Create(user.ToTransfer()).Error
}

func (r *userRepository) GetByUsername(username string) (response *models.User, err error) {
	err = r.db.Find(&response, "username LIKE", username).Error
	return
}


// C = conhecimento
// E = estrategia
// P = preparação
// A = Aplicação