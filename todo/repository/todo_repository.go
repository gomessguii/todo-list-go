package repository

import (
	"github.com/gomessguii/todo-list-go/todo/models"
	"gorm.io/gorm"
)

type ITodoRepository interface {
	Get(id int) (*models.Todo, error)
	CreateTodo(todo *models.Todo) (error)
	UpdateTodo(todo *models.Todo) (*models.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) ITodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) CreateTodo(todo *models.Todo) error {
	return r.db.Create(todo.ToTransfer()).Error
}

func (r *todoRepository) Get(id int) (response *models.Todo, err error) {
	err = r.db.Find(&response, id).Error
	return
}

func (r *todoRepository) UpdateTodo(todo *models.Todo) (*models.Todo, error) {
	foundTodo, err := r.Get(todo.ToTransfer().ID)
	if err != nil {
		return nil, err
	}
	err = r.db.Updates(foundTodo).Error
	if err != nil {
		return nil, err
	}
	return foundTodo, err
}

