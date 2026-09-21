package controller

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/dto"
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/model"
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/service"
	"github.com/your-org/ums-bff-service-customer-360/pkg/response"
	"github.com/your-org/ums-bff-service-customer-360/pkg/validator"
)

type OutletController struct {
	service service.OutletService
}

func NewOutletController(svc service.OutletService) *OutletController {
	return &OutletController{service: svc}
}

func (c *OutletController) GetAll(echoCtx echo.Context) error {
	limit, offset, err := validator.ParsePagination(echoCtx.QueryParams())
	if err != nil {
		return echoCtx.JSON(400, response.Error(err.Error()))
	}

	outlets, err := c.service.GetAll(limit, offset)
	if err != nil {
		return echoCtx.JSON(500, response.Error(err.Error()))
	}

	return echoCtx.JSON(200, response.Success(outlets, "outlets retrieved"))
}

func (c *OutletController) GetByID(echoCtx echo.Context) error {
	id, err := strconv.ParseInt(echoCtx.Param("id"), 10, 64)
	if err != nil {
		return echoCtx.JSON(400, response.Error("invalid id"))
	}

	outlet, err := c.service.GetByID(id)
	if err != nil {
		return echoCtx.JSON(500, response.Error(err.Error()))
	}
	if outlet == nil {
		return echoCtx.JSON(404, response.Error("outlet not found"))
	}

	return echoCtx.JSON(200, response.Success(outlet, "outlet retrieved"))
}

func (c *OutletController) Create(echoCtx echo.Context) error {
	var req dto.CreateOutletRequest
	if err := echoCtx.Bind(&req); err != nil {
		return echoCtx.JSON(400, response.Error(err.Error()))
	}

	if err := validator.Required(req.Name, "name"); err != nil {
		return echoCtx.JSON(400, response.Error(err.Error()))
	}
	if err := validator.Required(req.Code, "code"); err != nil {
		return echoCtx.JSON(400, response.Error(err.Error()))
	}

	outlet := &model.Outlet{
		Name:     req.Name,
		Code:     req.Code,
		Address:  req.Address,
		City:     req.City,
		Province: req.Province,
		IsActive: true,
	}

	if err := c.service.Create(outlet); err != nil {
		return echoCtx.JSON(500, response.Error(err.Error()))
	}

	return echoCtx.JSON(201, response.Success(outlet, "outlet created"))
}

func (c *OutletController) Update(echoCtx echo.Context) error {
	id, err := strconv.ParseInt(echoCtx.Param("id"), 10, 64)
	if err != nil {
		return echoCtx.JSON(400, response.Error("invalid id"))
	}

	var req dto.UpdateOutletRequest
	if err := echoCtx.Bind(&req); err != nil {
		return echoCtx.JSON(400, response.Error(err.Error()))
	}

	outlet, err := c.service.GetByID(id)
	if err != nil {
		return echoCtx.JSON(500, response.Error(err.Error()))
	}
	if outlet == nil {
		return echoCtx.JSON(404, response.Error("outlet not found"))
	}

	outlet.Name = req.Name
	outlet.Code = req.Code
	outlet.Address = req.Address
	outlet.City = req.City
	outlet.Province = req.Province

	if err := c.service.Update(outlet); err != nil {
		return echoCtx.JSON(500, response.Error(err.Error()))
	}

	return echoCtx.JSON(200, response.Success(outlet, "outlet updated"))
}

func (c *OutletController) Delete(echoCtx echo.Context) error {
	id, err := strconv.ParseInt(echoCtx.Param("id"), 10, 64)
	if err != nil {
		return echoCtx.JSON(400, response.Error("invalid id"))
	}

	if err := c.service.Delete(id); err != nil {
		return echoCtx.JSON(500, response.Error(err.Error()))
	}

	return echoCtx.JSON(200, response.Success(nil, "outlet deleted"))
}