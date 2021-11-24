package Interface

import "github.com/labstack/echo"

type UserAPI interface {
	SearchUsersHandler(c echo.Context) error
	GetUserHandler(c echo.Context) error
	UpdateUserHandler(c echo.Context) error
	CreateUserHandler(c echo.Context) error
	DeleteUserHandler(c echo.Context) error
}
