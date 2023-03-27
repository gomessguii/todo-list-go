package models

type Builder interface {
	WithID(id int) Builder
	WithTask(task string) Builder
	WithCompleted(completed bool) Builder
	Build() *Todo
}

type builder struct {
	todo *Todo
}

func NewBuilder() Builder {
	return &builder{todo: &Todo{}}
}

func (b *builder) WithID(id int) Builder {
	b.todo.id = id
	return b
}

func (b *builder) WithTask(task string) Builder {
	b.todo.task = task
	return b
}

func (b *builder) WithCompleted(completed bool) Builder {
	b.todo.completed = completed
	return b
}

func (b *builder) Build() *Todo {
	return b.todo
}