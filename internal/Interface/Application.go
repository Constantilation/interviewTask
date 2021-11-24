package Interface

import (
	"interviewTask/internal/User"
)

type UserApplication interface {
	SearchUsers() error
	GetUser(id int) error
	UpdateUser(userUpdate *User.UpdateUserRequest) error
	CreateUser(userGet *User.CreateUserRequest) error
	DeleteUser(id int) error
}
