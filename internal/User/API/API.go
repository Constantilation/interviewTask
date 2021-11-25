package UserAPI

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
	"interviewTask/internal/Interface"
	errPkg "interviewTask/internal/Middleware/Error"
	"interviewTask/internal/User"
	"net/http"
	"strconv"
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
	ctx := c.Request().Context()

	checkError := &errPkg.CheckError{
		Logger: A.Logger,
	}

	userArray, err := A.Application.SearchUsers(ctx)
	if err != nil {
		return c.JSON(checkError.CheckErrorsUser(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, userArray)
}

func (A API) GetUserHandler(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errPkg.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	user, err := A.Application.GetUser(ctx, id)

	checkError := &errPkg.CheckError{
		Logger: A.Logger,
	}

	if err != nil {
		return c.JSON(checkError.CheckErrorsUser(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (A API) UpdateUserHandler(c echo.Context) error {
	var updateUserStruct User.UpdateUserRequest
	err := c.Bind(&updateUserStruct)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	checkError := &errPkg.CheckError{
		Logger: A.Logger,
	}

	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(checkError.CheckErrorsUser(err), ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()

	err = A.Application.UpdateUser(ctx, &updateUserStruct, idP)
	if err != nil {
		return c.JSON(checkError.CheckErrorsUser(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func isRequestValid(m *User.CreateUserRequest) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}

	return true, nil
}

// CreateUserHandler will initialize the users/ resources endpoint
func (A API) CreateUserHandler(c echo.Context) error {
	var createUserStruct User.CreateUserRequest
	err := c.Bind(&createUserStruct)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&createUserStruct); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err, id := A.Application.CreateUser(ctx, &createUserStruct)

	checkError := &errPkg.CheckError{
		Logger: A.Logger,
	}

	if err != nil {
		return c.JSON(checkError.CheckErrorsUser(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, id)
}

func (A API) DeleteUserHandler(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, errPkg.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = A.Application.DeleteUser(ctx, id)

	checkError := &errPkg.CheckError{
		Logger: A.Logger,
	}

	if err != nil {
		return c.JSON(checkError.CheckErrorsUser(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

// NewUserHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, handler Interface.UserAPI, url string) {
	ug := e.Group(url)
	ug.GET("/", handler.SearchUsersHandler)
	ug.POST("/", handler.CreateUserHandler)

	ug2 := ug.Group("/:id")
	ug2.GET("/", handler.GetUserHandler)
	ug2.PATCH("/", handler.UpdateUserHandler)
	ug2.DELETE("/", handler.DeleteUserHandler)

}
