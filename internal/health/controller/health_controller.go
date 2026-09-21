package controller

import (
	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (c *HealthController) Check(echoCtx echo.Context) error {
	return echoCtx.JSON(200, map[string]string{
		"status": "ok",
	})
}
