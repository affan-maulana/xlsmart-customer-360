package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/vqm/dto"
	"github.com/your-org/ums-bff-service-customer-360/pkg/response"
)

type VQMController struct {
	client *client.Client
}

func NewVQMController(c *client.Client) *VQMController {
	return &VQMController{client: c}
}

func (ctrl *VQMController) GetProfile(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	res := ctrl.client.Get("/vqm/widget/main-profile/get-profile/" + msisdn)
	if !res.Success {
		return ctx.JSON(http.StatusBadGateway, res)
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to marshal response data"))
	}

	var raw dto.RawGetProfileData
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to parse third-party response"))
	}

	addrParts := []string{}
	if raw.Address.HomeAddress.StreetName != "" {
		addrParts = append(addrParts, raw.Address.HomeAddress.StreetName)
	}
	if raw.Address.HomeAddress.City != "" && raw.Address.HomeAddress.City != "Anonymous" {
		addrParts = append(addrParts, raw.Address.HomeAddress.City)
	}
	if raw.Address.HomeAddress.Country != "" {
		addrParts = append(addrParts, raw.Address.HomeAddress.Country)
	}
	address := strings.Join(addrParts, ", ")

	formattedBirthDate := ""
	if len(raw.BirthDate) >= 10 {
		formattedBirthDate = raw.BirthDate[:10]
	}

	mobileBalanceNumeric, _ := strconv.ParseFloat(raw.MainBalance.MobileBalance, 64)
	dealerBalanceNumeric, _ := strconv.ParseFloat(raw.MainBalance.DealerBalance, 64)

	mapped := dto.GetProfileResponse{
		Msisdn:                  raw.Msisdn,
		FullName:                raw.FullName.GivenName,
		Gender:                  raw.Gender,
		BirthDate:               raw.BirthDate,
		FormattedBirthDate:      formattedBirthDate,
		Address:                 address,
		ContactRole:             raw.AdditionalInformation.ContactRole,
		HomePhone:               raw.AdditionalInformation.HomePhone,
		MaritalStatus:           raw.AdditionalInformation.MaritalStatus,
		Occupation:              raw.AdditionalInformation.Occupation,
		Income:                  raw.AdditionalInformation.Income,
		Religion:                raw.AdditionalInformation.Religion,
		SubscriberStatusMapped:  "Active",
		CustomerId:              raw.CustomerId,
		CustomerRank:            raw.CustomerRank,
		CustomerSubType:         raw.CustomerSubType,
		CustomerType:            raw.CustomerType,
		IndividualId:            raw.IndividualId,
		InitialRegistrationDate: raw.InitialRegistrationDate,
		MobileBalance:           raw.MainBalance.MobileBalance,
		MobileBalanceNumeric:    mobileBalanceNumeric,
		DealerBalance:           raw.MainBalance.DealerBalance,
		DealerBalanceNumeric:    dealerBalanceNumeric,
		PaymentType:             raw.PaymentType,
		SubscriberType:          raw.SubscriberType,
		IsCustomerCorporate:     raw.IsCustomerCorporate,
		FirstEventDate:          raw.FirstEventDate,
		IccId:                   raw.IccId,
		CustomerName:            raw.CustomerName,
		BillCycle:               raw.BillCycle,
		PaymentMethod:           raw.OpGetPaymentMethodDetails.PaymentMethod,
		CcNumber:                raw.OpGetPaymentMethodDetails.CcNumber,
		CcType:                  raw.OpGetPaymentMethodDetails.CcType,
		CcExpDate:               raw.OpGetPaymentMethodDetails.CcExpDate,
	}

	return ctx.JSON(http.StatusOK, response.Success(mapped, "success"))
}

func (ctrl *VQMController) GetCustomerStatus(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	res := ctrl.client.Get("/vqm/widget/main-profile/get-customer-status/" + msisdn)
	if !res.Success {
		return ctx.JSON(http.StatusBadGateway, res)
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to marshal response data"))
	}

	var mapped dto.GetCustomerStatusResponse
	if err := json.Unmarshal(rawBytes, &mapped); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to parse third-party response"))
	}

	return ctx.JSON(http.StatusOK, response.Success(mapped, "success"))
}

func (ctrl *VQMController) GetPackageDetails(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	res := ctrl.client.Get("/vqm/page/package-details/" + msisdn)
	if !res.Success {
		return ctx.JSON(http.StatusBadGateway, res)
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to marshal response data"))
	}

	var mapped dto.GetPackageDetailsResponse
	if err := json.Unmarshal(rawBytes, &mapped); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to parse third-party response"))
	}

	return ctx.JSON(http.StatusOK, response.Success(mapped, "success"))
}

func (ctrl *VQMController) GetHLRDetails(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	res := ctrl.client.Get("/vqm/page/hlr-details/" + msisdn)
	if !res.Success {
		return ctx.JSON(http.StatusBadGateway, res)
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to marshal response data"))
	}

	var mapped dto.GetHLRDetailsResponse
	if err := json.Unmarshal(rawBytes, &mapped); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to parse third-party response"))
	}

	return ctx.JSON(http.StatusOK, response.Success(mapped, "success"))
}
