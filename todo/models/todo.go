package models

type Todo struct {
	id int
	task string
	completed bool
}

func (u *Todo) CompleteTodo() {
	u.completed = true
}

func (u *Todo) ToTransfer() *TodoTransfer {
	return &TodoTransfer{ID: u.id, Task: u.task, Completed: u.completed}
}

type TodoTransfer struct {
	ID int `json:"id"`
	Task string `json:"task"`
	Completed bool `json:"completed"`
}