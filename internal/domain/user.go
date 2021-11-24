package domain

import (
	"context"
	"time"
)

// User struct
type User struct {
CreatedAt   time.Time `json:"created_at"`
DisplayName string    `json:"display_name"`
Email       string    `json:"email"`
}

// UserUsecase represent the article's usecases
type UserUsecase interface {
	SearchUsers(ctx context.Context, cursor string, num int64) ([]User, string, error)
	GetUser(ctx context.Context, id int64) (User, error)
	UpdateUser(ctx context.Context, ar *User) error
	CreateUser(context.Context, *User) error
	DeleteUser(ctx context.Context, id int64) error
}

// UserRepository represent the user's repository contract
type UserRepository interface {
	SearchUsers(ctx context.Context, cursor string, num int64) ([]User, string, error)
	GetUser(ctx context.Context, id int64) (User, error)
	UpdateUser(ctx context.Context, ar *User) error
	CreateUser(context.Context, *User) error
	DeleteUser(ctx context.Context, id int64) error
}