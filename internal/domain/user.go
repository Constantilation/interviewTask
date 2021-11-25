package domain

import (
	"time"
)

// User struct
type (
	User struct {
		CreatedAt   time.Time `json:"created_at"`
		DisplayName string    `json:"display_name"`
		Email       string    `json:"email"`
	}
	UserList  map[string]User
	UserStore struct {
		CustID    string   `json:"custid1"`
		Increment int      `json:"increment"`
		List      UserList `json:"list"`
	}
)

// ID implementation of func for database library
func (u UserStore) ID() (jsonField string, value interface{}) {
	value = u.CustID
	jsonField = "custid2"
	return
}
