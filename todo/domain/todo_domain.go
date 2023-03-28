package domain

import (
	"fmt"

	"github.com/gomessguii/todo-list-go/todo/models"
)

type ITodoDomain interface {
	GetTodo(id int) (*models.Todo, error)
	CreateTodo(todo *models.Todo) error
	CompleteTodo(todoID int) error
}

type todoDomain struct {
	userDomain     IUserDomain
	todoRepository ITodoRepository
}

func New(userDomain IUserDomain, todoRepo ITodoRepository) ITodoDomain {
	return &todoDomain{userDomain: userDomain, todoRepository: todoRepo}
}


func (d *todoDomain) CompleteTodo(todoID int) error {
	todo, err := d.todoRepository.Get(todoID)
	if err != nil {
		return err
	}
	todo.CompleteTodo()
	_, err = d.todoRepository.UpdateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (d *todoDomain) CreateTodo(todo *models.Todo) error {
	if !todo.IsValid() {
		return fmt.Errorf("invalid entity")
	}
	err := d.todoRepository.CreateTodo(todo)
	if err != nil {
		return err
	}
	return nil
}

func (d *todoDomain) GetTodo(id int) (*models.Todo, error) {
	todo, err := d.todoRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

