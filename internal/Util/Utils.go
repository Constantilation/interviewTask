package Util

import (
	errors "interviewTask/internal/Middleware/Error"
	"strconv"
)

// InterfaceConvertInt just some type of utils func, didn't use it in project
func InterfaceConvertInt(value interface{}) (int, error) {
	var intConvert int
	var errorConvert error
	switch value.(type) {
	case string:
		intConvert, errorConvert = strconv.Atoi(value.(string))
		if errorConvert != nil {
			return errors.IntNil, &errors.Errors{
				Alias: errors.ErrAtoi,
			}
		}
		return intConvert, nil
	case int:
		intConvert = value.(int)
		return intConvert, nil
	default:
		return errors.IntNil, &errors.Errors{
			Alias: errors.ErrNotStringAndInt,
		}
	}
}
