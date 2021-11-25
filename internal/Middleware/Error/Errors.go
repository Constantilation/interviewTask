package Error

import "errors"

type MultiLogger interface {
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Sync() error
}

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")
)

const (
	ErrMarshal         = "marshaling in json"
	ErrCheck           = "err check"
	IntNil             = 0
	ErrNotStringAndInt = "expected type string or int"
	ErrAtoi            = "func Atoi convert string in int"
)

type ResultError struct {
	Status  int    `json:"status"`
	Explain string `json:"explain,omitempty"`
}

type Errors struct {
	Alias string
	Text  string
}

func (e *Errors) Error() string {
	return e.Alias
}

type CheckError struct {
	Logger MultiLogger
}

