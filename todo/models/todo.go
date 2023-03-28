package models

type Todo struct {
	id   int
	task string
	done bool
}

func (t *Todo) IsValid() bool {
	if t.task == "" {
		return false
	}
	return true
}

func (u *Todo) CompleteTodo() {
	u.done = true
}

func (u *Todo) ToTransfer() *TodoTransfer {
	return &TodoTransfer{ID: u.id, Task: u.task, Done: &u.done}
}

type TodoTransfer struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done *bool  `json:"completed"`
}

func (u *TodoTransfer) ToModel() *Todo {
	return &Todo{
		id:   u.ID,
		task: u.Task,
		done: *u.Done,
	}
}
