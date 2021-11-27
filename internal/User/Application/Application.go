package UserApplication

import (
	"context"
	"interviewTask/internal/Interface"
	"interviewTask/internal/User"
	"interviewTask/internal/domain"
	"time"
)

// Application struct for use cases level
type Application struct {
	Store Interface.UserStore
}

// SearchUsers represent the application method to search users
func (a Application) SearchUsers(ctx context.Context) (domain.UserList, error) {
	res, err := a.Store.SearchUsers(ctx)
	if err != nil {
		return nil, err
	}

	return res, err
}

// GetUser represent the application method to get user by id
func (a Application) GetUser(ctx context.Context, id int64) (domain.User, error) {
	res, err := a.Store.GetUser(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

// UpdateUser represent the application method to update user info
func (a Application) UpdateUser(ctx context.Context, u *User.UpdateUserRequest, id int) error {
	err := a.Store.UpdateUser(ctx, u, id)
	if err != nil {
		return err
	}

	return err
}

// CreateUser represent the application method to create new user
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

// DeleteUser represent the application method to delete user from store
func (a Application) DeleteUser(ctx context.Context, id int64) error {
	err := a.Store.DeleteUser(ctx, id)
	if err != nil {
		return err
	}

	return err
}
