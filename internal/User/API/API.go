package UserAPI

import (
	"github.com/labstack/echo"
	"interviewTask/internal/Interface"
	errPkg "interviewTask/internal/MyError"
)

type API struct {
	Application Interface.UserApplication
	Logger      errPkg.MultiLogger
}

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// SearchUsersHandler represent the httphandler for users
func (A API) SearchUsersHandler(c echo.Context) error {
	panic("implement me")
}

func (A API) GetUserHandler(c echo.Context) error {
	panic("implement me")
}

func (A API) UpdateUserHandler(c echo.Context) error {
	panic("implement me")
}
// CreateUserHandler will initialize the users/ resources endpoint
func (A API) CreateUserHandler(c echo.Context) error {
	panic("implement me")
}

func (A API) DeleteUserHandler(c echo.Context) error {
	panic("implement me")
}

// NewUserHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, handler Interface.UserAPI, url string) {
	ug := e.Group(url)
	ug.GET("/", handler.SearchUsersHandler)
	ug.POST("/", handler.CreateUserHandler)

	ug2 := ug.Group("/{id}")
	ug2.GET("/", handler.GetUserHandler)
	ug2.PATCH("/", handler.UpdateUserHandler)
	ug2.DELETE("/", handler.DeleteUserHandler)

}


