package domain

type UserInterface interface {
	Create(user *User) error
	GetUsers() ([]*User, error)
	GetUserById(id int32) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int32) error
}
