package health

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/health/controller"
)

func RegisterRoutes(e *echo.Echo) {
	healthCtrl := controller.NewHealthController()
	e.GET("/health", healthCtrl.Check)
}