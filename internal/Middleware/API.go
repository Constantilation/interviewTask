package Middleware

import (
	"github.com/labstack/echo"
	errPkg "interviewTask/internal/Middleware/Error"
	"math"
)

// GoMiddleware represent the data-struct for middleware
type GoMiddleware struct {
	// another stuff , may be needed by middleware
}

type InfoMiddleware struct {
	Logger      errPkg.MultiLogger
	ReqId       int
}

func (m *InfoMiddleware) LogURL(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if m.ReqId == math.MaxInt {
			m.ReqId = 0
		}
		m.ReqId++
		m.Logger.Infof("Method: %s, URL: %s, from: %s - %d, requestId: %d", string(c.Request().Method), c.Request().URL, c.Request().Host, c.Response().Status, m.ReqId)
		return next(c)
	}
}

// CORS will handle the CORS middleware
func (m *GoMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}

// InitMiddleware initialize the middleware
func InitMiddleware() *GoMiddleware {
	return &GoMiddleware{}
}

