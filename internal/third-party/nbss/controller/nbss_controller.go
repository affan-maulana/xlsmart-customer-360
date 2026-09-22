package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/nbss/dto"
	"github.com/your-org/ums-bff-service-customer-360/pkg/response"
)

type NBSSController struct {
	client *client.Client
}

func NewNBSSController(c *client.Client) *NBSSController {
	return &NBSSController{client: c}
}

func (ctrl *NBSSController) GetDeviceSpecs(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	res := ctrl.client.Get("/nbss/device-specs/" + msisdn)
	if !res.Success {
		return ctx.JSON(http.StatusBadGateway, res)
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to marshal response data"))
	}

	var raw dto.RawDeviceSpecsResponse
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to parse third-party response"))
	}

	mapped := dto.GetDeviceSpecsResponse{
		Manufacturer: raw.Data.TCR_NAME,
		ModelName:    raw.Data.TCR_MODEL,
		OsName:       raw.Data.OS_NAME,
		OsVersion:    raw.Data.OS_VERSION,
		Imei:         raw.Data.IMEI,
	}

	return ctx.JSON(http.StatusOK, response.Success(mapped, "success"))
}
