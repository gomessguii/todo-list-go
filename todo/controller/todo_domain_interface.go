package controller

import "github.com/gomessguii/todo-list-go/todo/models"

type ITodoDomain interface {
	GetTodo(id int) (*models.Todo, error)
	CreateTodo(todo *models.Todo) error
	CompleteTodo(todoID int) error
}