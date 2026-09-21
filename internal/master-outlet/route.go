package masteroutlet

import (
	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/controller"
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/service"
	"github.com/your-org/ums-bff-service-customer-360/internal/middleware"
)

func RegisterRoutes(e *echo.Echo, svc service.OutletService) {
	outletCtrl := controller.NewOutletController(svc)
	grp := e.Group("/api/v1/outlets", middleware.Auth())
	grp.GET("", outletCtrl.GetAll)
	grp.GET("/:id", outletCtrl.GetByID)
	grp.POST("", outletCtrl.Create)
	grp.PUT("/:id", outletCtrl.Update)
	grp.DELETE("/:id", outletCtrl.Delete)
}