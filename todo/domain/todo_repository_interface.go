package domain

type ITodoRepository interface {
	CreateTask(task string) (string, error)
	UpdateTask(task string) (string, error)
}