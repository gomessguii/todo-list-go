package models

type Builder interface {
	WithID(id int) Builder
	WithUsername(username string) Builder
	WithPassword(password string) Builder
	Build() *User
}

type builder struct {
	user *User
}

func NewBuilder() Builder {
	return &builder{user: &User{}}
}

func (b *builder) WithID(id int) Builder {
	b.user.id = id
	return b
}

func (b *builder) WithPassword(password string) Builder {
	b.user.password = password
	return b
}

func (b *builder) WithUsername(username string) Builder {
	b.user.username = username
	return b
}

func (b *builder) Build() *User {
	return b.user
}