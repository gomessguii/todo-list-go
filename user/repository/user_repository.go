package repository


type IUserRepository interface {
	GetByUsername(username string) (string, error)
	CreateUser(username, password string) error
}

func New() IUserRepository {
	panic("implement me")
}