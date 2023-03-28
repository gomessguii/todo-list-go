package models

type User struct {
	id       int
	username string
	password string
}

func (u *User) ToTransfer() *UserTransfer {
	return &UserTransfer{
		ID:       u.id,
		Username: u.username,
		Password: u.password,
	}
}

func (u *User) CheckPassword(password string) bool {
	return u.password == password
}

type UserTransfer struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserTransfer)ToModel() *User {
	return &User{
		id: u.ID, 
		username: u.Username, 
		password: u.Password,
	}
}
