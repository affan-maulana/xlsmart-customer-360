package vqm

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/vqm/controller"
)

func RegisterRoutes(e *echo.Echo, c *client.Client) {
	ctrl := controller.NewVQMController(c)
	// grp := e.Group("/api/v1/third-party/vqm", middleware.Auth())
	grp := e.Group("/api/v1/third-party/vqm")

	grp.GET("/get-profile/:msisdn", ctrl.GetProfile)
	grp.GET("/get-customer-status/:msisdn", ctrl.GetCustomerStatus)
	grp.GET("/package-details/:msisdn", ctrl.GetPackageDetails)
	grp.GET("/hlr-details/:msisdn", ctrl.GetHLRDetails)
}
