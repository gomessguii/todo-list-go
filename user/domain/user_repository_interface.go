package domain

type IUserRepository interface {
	GetByUsername(username string) (string, error)
	CreateUser(username, password string) error
}