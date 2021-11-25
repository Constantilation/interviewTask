package UserApplication

import (
	"context"
	"interviewTask/internal/Interface"
	"interviewTask/internal/User"
	"interviewTask/internal/domain"
	"time"
)

type Application struct {
	Store          Interface.UserStore
	contextTimeout time.Duration
}

func (a Application) SearchUsers(ctx context.Context) (domain.UserList, error) {
	ctxApp, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	res, err := a.Store.SearchUsers(ctxApp)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (a Application) GetUser(ctx context.Context, id int64) (domain.User, error) {
	res, err := a.Store.GetUser(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (a Application) UpdateUser(ctx context.Context, u *User.UpdateUserRequest, id int) error {
	err := a.Store.UpdateUser(ctx, u, id)
	if err != nil {
		return err
	}

	return err
}

func (a Application) CreateUser(ctx context.Context, request *User.CreateUserRequest) (error, map[string]interface{}) {
	user := domain.User{
		Email:       request.Email,
		DisplayName: request.DisplayName,
		CreatedAt:   time.Now(),
	}

	err, id := a.Store.CreateUser(ctx, user)
	if err != nil {
		return err, nil
	}

	return nil, map[string]interface{}{
		"user_id": id,
	}
}

func (a Application) DeleteUser(ctx context.Context, id int64) error {
	err := a.Store.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return err
}
