package UserApplication

import (
	"interviewTask/internal/Interface"
	"interviewTask/internal/User"
)

type Application struct {
	Store Interface.UserApplication
}

func (a Application) SearchUsers() error {
	panic("implement me")
}

func (a Application) GetUser(id int) error {
	panic("implement me")
}

func (a Application) UpdateUser(userUpdate *User.UpdateUserRequest) error {
	panic("implement me")
}

func (a Application) CreateUser(userGet *User.CreateUserRequest) error {
	panic("implement me")
}

func (a Application) DeleteUser(id int) error {
	panic("implement me")
}
