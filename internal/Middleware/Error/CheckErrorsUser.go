package Error

import (
	"net/http"
)

func (c CheckError) CheckErrorsUser(errIn error) (int) {
	if errIn == nil {
		return http.StatusOK
	}

	c.Logger.Errorf("",errIn.Error(), errIn)
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
