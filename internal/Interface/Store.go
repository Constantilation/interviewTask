package Interface

import (
	"context"
	"github.com/sonyarouje/simdb"
	"interviewTask/internal/User"
	"interviewTask/internal/domain"
)

type ConnectionInterface interface {
	Open(entity simdb.Entity) *simdb.Driver
	Errors() []error
	Insert(entity simdb.Entity) (err error)
	Where(key string, cond string, val interface{}) *simdb.Driver
	Get() *simdb.Driver
	First() *simdb.Driver
	Raw() interface{}
	RawArray() []interface{}
	AsEntity(output interface{}) (err error)
	Update(entity simdb.Entity) (err error)
	Upsert(entity simdb.Entity) (err error)
	Delete(entity simdb.Entity) (err error)
}

type UserStore interface {
	SearchUsers(ctx context.Context) (domain.UserList, error)
	GetUser(ctx context.Context, id int64) (domain.User, error)
	UpdateUser(ctx context.Context, u *User.UpdateUserRequest, id int) error
	CreateUser(context.Context, domain.User) (error, int)
	DeleteUser(ctx context.Context, id int64) error
}
