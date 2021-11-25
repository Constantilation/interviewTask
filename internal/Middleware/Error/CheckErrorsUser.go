package Error

import (
	"net/http"
)

// CheckErrorsUser method to check some errors, can be used as middleware, also can be expanded with a lot of errors
func (c CheckError) CheckErrorsUser(errIn error) int {
	if errIn == nil {
		return http.StatusOK
	}

	c.Logger.Errorf("", errIn.Error(), errIn)
	switch errIn {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
