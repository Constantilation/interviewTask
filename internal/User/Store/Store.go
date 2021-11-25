package UserStore

import (
	"context"
	"interviewTask/internal/Interface"
	"interviewTask/internal/User"
	"interviewTask/internal/domain"
	"strconv"
)

type Store struct {
	Conn Interface.ConnectionInterface
}

func (s Store) SearchUsers(ctx context.Context) (domain.UserList, error) {
	var users domain.UserStore
	err := s.Conn.Open(domain.UserStore{}).Where("custid1", "=", "CUSTID1").First().AsEntity(&users)
	if err != nil {
		return nil, err
	}

	return users.List, nil
}

func (s Store) GetUser(ctx context.Context, id int64) (domain.User, error) {
	var users domain.UserStore
	err := s.Conn.Open(domain.UserStore{}).Where("custid1", "=", "CUSTID1").First().AsEntity(&users)
	if err != nil {
		return users.List["0"], err
	}

	if element, ok := users.List[strconv.Itoa(int(id))]; ok {
		return element, nil
	}

	return users.List["0"], err
}

func (s Store) UpdateUser(ctx context.Context, u *User.UpdateUserRequest, id int) error {
	var users domain.UserStore
	users.List = make(map[string]domain.User)
	err := s.Conn.Open(domain.UserStore{}).Where("custid1", "=", "CUSTID1").First().AsEntity(&users)
	if err != nil {
		return err
	}
	err = s.Conn.Delete(users)
	if err != nil {
		return err
	}

	if element, ok := users.List[strconv.Itoa(id)]; ok {
		element.DisplayName = u.DisplayName
		users.List[strconv.Itoa(id)] = element
	}

	err = s.Conn.Insert(users)
	if err != nil {
		return err
	}

	return nil
}

func (s Store) CreateUser(ctx context.Context, user domain.User) (error, int) {
	var users domain.UserStore
	users.List = make(map[string]domain.User)
	err := s.Conn.Open(domain.UserStore{}).Where("custid1", "=", "CUSTID1").First().AsEntity(&users)
	if err != nil {
		users.Increment++
		users.List[strconv.Itoa(users.Increment)] = user
		err = s.Conn.Insert(users)
		if err != nil {
			return err, 0
		}

		return err, users.Increment
	}
	err = s.Conn.Delete(users)
	if err != nil {
		return err, 0
	}

	users.Increment++
	users.List[strconv.Itoa(users.Increment)] = user
	err = s.Conn.Insert(users)
	if err != nil {
		return err, 0
	}

	return nil, users.Increment
}

func (s Store) DeleteUser(ctx context.Context, id int64) error {
	var users domain.UserStore
	users.List = make(map[string]domain.User)
	err := s.Conn.Open(domain.UserStore{}).Where("custid1", "=", "CUSTID1").First().AsEntity(&users)
	if err != nil {
		return err
	}

	err = s.Conn.Delete(users)
	if err != nil {
		return err
	}

	if _, ok := users.List[strconv.Itoa(int(id))]; ok {
		delete(users.List, strconv.Itoa(int(id)))
	}

	err = s.Conn.Insert(users)
	if err != nil {
		return err
	}

	return nil
}
