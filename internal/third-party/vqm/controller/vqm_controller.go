package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
)

type VQMController struct {
	client *client.Client
}

func NewVQMController(c *client.Client) *VQMController {
	return &VQMController{client: c}
}

func (ctrl *VQMController) GetProfile(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	resp := ctrl.client.Get("/vqm/widget/main-profile/get-profile/" + msisdn)
	return ctx.JSON(200, resp)
}

func (ctrl *VQMController) GetCustomerStatus(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	resp := ctrl.client.Get("/vqm/widget/main-profile/get-customer-status/" + msisdn)
	return ctx.JSON(200, resp)
}

func (ctrl *VQMController) GetPackageDetails(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	resp := ctrl.client.Get("/vqm/page/package-details/" + msisdn)
	return ctx.JSON(200, resp)
}

func (ctrl *VQMController) GetHLRDetails(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	resp := ctrl.client.Get("/vqm/page/hlr-details/" + msisdn)
	return ctx.JSON(200, resp)
}
