package thirdparty

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/config"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/nbss"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/vqm"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config) {
	vqmClient := client.NewClient(cfg.VQMBaseURL)
	nbssClient := client.NewClient(cfg.NBSSBaseURL)

	vqm.RegisterRoutes(e, vqmClient)
	nbss.RegisterRoutes(e, nbssClient)
}
