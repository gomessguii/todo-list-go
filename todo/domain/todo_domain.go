package domain

type ITodoDomain interface {
	CreateTask(task string) error
	CompleteTask(taskID int) error
}

func New(IUserDomain, ITodoRepository) ITodoDomain {
	panic("implement me")
}