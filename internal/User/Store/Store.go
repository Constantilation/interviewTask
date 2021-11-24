package UserStore

import (
	"interviewTask/internal/Interface"
	"interviewTask/internal/User"
)

type Store struct {
	Conn Interface.ConnectionInterface
}

func (w Store) SearchUsers() error {
	panic("implement me")
}

func (w Store) GetUser(id int) error {
	panic("implement me")
}

func (w Store) UpdateUser(userUpdate *User.UpdateUserRequest) error {
	panic("implement me")
}

func (w Store) CreateUser(userGet *User.CreateUserRequest) error {
	panic("implement me")
}

func (w Store) DeleteUser(id int) error {
	panic("implement me")
}

