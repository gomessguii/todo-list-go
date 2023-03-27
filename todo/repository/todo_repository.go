package repository

import (
	"github.com/gomessguii/todo-list-go/todo/models"
)

type ITodoRepository interface {
	Get(id int) (*models.Todo, error)
	CreateTodo(todo *models.Todo) (*models.Todo, error)
	UpdateTodo(todo *models.Todo) (*models.Todo, error)
}

func New() ITodoRepository {
	panic("implement me")
}
