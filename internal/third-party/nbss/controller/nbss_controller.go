package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
)

type NBSSController struct {
	client *client.Client
}

func NewNBSSController(c *client.Client) *NBSSController {
	return &NBSSController{client: c}
}

func (ctrl *NBSSController) GetDeviceSpecs(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	resp := ctrl.client.Get("/nbss/device-specs/" + msisdn)
	return ctx.JSON(200, resp)
}
