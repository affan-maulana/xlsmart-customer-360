package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
)

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := c.Response().Header().Get(echo.HeaderXRequestID)
			if reqID == "" {
				reqID = time.Now().Format("20060102150405")
				c.Response().Header().Set(echo.HeaderXRequestID, reqID)
			}
			c.Set("request_id", reqID)
			return next(c)
		}
	}
}