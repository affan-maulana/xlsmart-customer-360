package profile

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/profile/controller"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
)

func RegisterRoutes(e *echo.Echo, c *client.Client) {
	ctrl := controller.NewProfileController(c)
	// grp := e.Group("/api/v1/profile", middleware.Auth())
	grp := e.Group("/api/v1/profile")

	grp.GET("/msisdns", ctrl.GetMsisdnByNik)
	grp.GET("/msisdns/:msisdn", ctrl.GetMsisdnDetail)
}
