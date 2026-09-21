package nbss

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/nbss/controller"
)

func RegisterRoutes(e *echo.Echo, c *client.Client) {
	ctrl := controller.NewNBSSController(c)
	// grp := e.Group("/api/v1/third-party/nbss", middleware.Auth())
	grp := e.Group("/api/v1/third-party/nbss")

	grp.GET("/device-specs/:msisdn", ctrl.GetDeviceSpecs)
}
