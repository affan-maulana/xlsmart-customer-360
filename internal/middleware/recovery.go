package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Recovery() echo.MiddlewareFunc {
	return middleware.Recover()
}
