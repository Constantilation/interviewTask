package Interface

import (
	"context"
	"interviewTask/internal/User"
	"interviewTask/internal/domain"
)

// UserApplication implementation of user Application interface
type UserApplication interface {
	SearchUsers(ctx context.Context) (domain.UserList, error)
	GetUser(ctx context.Context, id int64) (domain.User, error)
	UpdateUser(ctx context.Context, u *User.UpdateUserRequest, id int) error
	CreateUser(context.Context, *User.CreateUserRequest) (error, map[string]interface{})
	DeleteUser(ctx context.Context, id int64) error
}
