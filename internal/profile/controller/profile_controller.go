package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/your-org/ums-bff-service-customer-360/internal/profile/dto"
	"github.com/your-org/ums-bff-service-customer-360/internal/third-party/client"
	"github.com/your-org/ums-bff-service-customer-360/pkg/logger"
	"github.com/your-org/ums-bff-service-customer-360/pkg/response"
)

type ProfileController struct {
	client *client.Client
}

func NewProfileController(c *client.Client) *ProfileController {
	return &ProfileController{client: c}
}

func (ctrl *ProfileController) GetMsisdnByNik(ctx echo.Context) error {
	nik := ctx.QueryParam("nik")
	if nik == "" {
		return ctx.JSON(http.StatusBadRequest, response.Error("nik is required"))
	}

	res := ctrl.client.Get("/nbss/getMsisdnByNik?nik=" + nik)
	if !res.Success {
		return ctx.JSON(http.StatusBadGateway, res)
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to marshal response data"))
	}

	var raw dto.RawMsisdnByNikResponse
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.Error("failed to parse third-party response"))
	}

	mapped := dto.GetMsisdnByNikResponse{
		Data:  raw.Data.Data,
		Count: raw.Data.RecordCount,
	}

	return ctx.JSON(http.StatusOK, response.Success(mapped, "success"))
}

func (ctrl *ProfileController) GetMsisdnDetail(ctx echo.Context) error {
	msisdn := ctx.Param("msisdn")
	if msisdn == "" {
		return ctx.JSON(http.StatusBadRequest, response.Error("msisdn is required"))
	}

	out := &dto.MsisdnDetailResponse{Msisdn: msisdn}

	ctrl.mergeGetProfile(msisdn, out)
	ctrl.mergeCustomerStatus(msisdn, out)
	ctrl.mergePackageDetails(msisdn, out)
	ctrl.mergeDeviceSpecs(msisdn, out)

	return ctx.JSON(http.StatusOK, response.Success(out, "success"))
}

func (ctrl *ProfileController) mergeGetProfile(msisdn string, out *dto.MsisdnDetailResponse) {
	res := ctrl.client.Get("/vqm/widget/main-profile/get-profile/" + msisdn)
	if !res.Success {
		logger.Error.Printf("profile: get-profile upstream failed for %s: %v", msisdn, res.Message)
		out.BirthDate = "-"
		out.MaritalStatus = "-"
		out.Gender = "-"
		out.Income = "-"
		out.JobTitle = "-"
		out.MothersMaidenName = "-"
		out.Religion = "-"
		out.DisplayCustomerType = "-"
		out.MobileBalance = "-"
		out.IndividualId = "-"
		out.IccId = "-"
		out.SpecialStatus = "-"
		out.GraceEndDate = "-"
		out.FirstEventDate = "-"
		out.ContactRole = "-"
		out.PaymentType = "-"
		return
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		logger.Error.Printf("profile: get-profile marshal failed for %s: %v", msisdn, err)
		out.BirthDate = "-"
		return
	}

	var raw dto.RawGetProfileData
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		logger.Error.Printf("profile: get-profile unmarshal failed for %s: %v", msisdn, err)
		out.BirthDate = "-"
		return
	}

	fullName := raw.FullName
	address := raw.Address
	pricePlan := raw.PricePlan

	out.BirthDate = raw.BirthDate
	out.FullName = &fullName
	out.MaritalStatus = raw.AdditionalInformation.MaritalStatus
	out.Gender = raw.Gender
	out.Address = &address
	out.Income = raw.AdditionalInformation.Income
	out.JobTitle = raw.AdditionalInformation.JobTitle
	out.MothersMaidenName = raw.AdditionalInformation.MothersMaidenName
	out.Religion = raw.AdditionalInformation.Religion
	out.PricePlan = &pricePlan
	out.DisplayCustomerType = raw.DisplayCustomerType
	out.MobileBalance = raw.MainBalance.MobileBalance
	out.IndividualId = raw.IndividualId
	out.IccId = raw.IccId
	out.SpecialStatus = raw.AdditionalInformation.SpecialStatus
	out.GraceEndDate = raw.GraceEndDate
	out.FirstEventDate = raw.FirstEventDate
	out.ContactRole = raw.AdditionalInformation.ContactRole
	out.PaymentType = raw.PaymentType
}

func (ctrl *ProfileController) mergeCustomerStatus(msisdn string, out *dto.MsisdnDetailResponse) {
	res := ctrl.client.Get("/vqm/widget/main-profile/get-customer-status/" + msisdn)
	if !res.Success {
		logger.Error.Printf("profile: get-customer-status upstream failed for %s: %v", msisdn, res.Message)
		out.Nik = "-"
		out.ActivationDate = "-"
		out.DukcapilStatus = "-"
		return
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		logger.Error.Printf("profile: get-customer-status marshal failed for %s: %v", msisdn, err)
		out.Nik = "-"
		return
	}

	var raw dto.RawCustomerStatusData
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		logger.Error.Printf("profile: get-customer-status unmarshal failed for %s: %v", msisdn, err)
		out.Nik = "-"
		return
	}

	out.Nik = raw.Nik
	out.ActivationDate = raw.ActivationDate
	out.DukcapilStatus = raw.DukcapilStatus
}

func (ctrl *ProfileController) mergePackageDetails(msisdn string, out *dto.MsisdnDetailResponse) {
	res := ctrl.client.Get("/vqm/page/package-details/" + msisdn)
	if !res.Success {
		logger.Error.Printf("profile: package-details upstream failed for %s: %v", msisdn, res.Message)
		out.PackageDetail = nil
		return
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		logger.Error.Printf("profile: package-details marshal failed for %s: %v", msisdn, err)
		out.PackageDetail = nil
		return
	}

	var raw dto.RawPackageDetailsData
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		logger.Error.Printf("profile: package-details unmarshal failed for %s: %v", msisdn, err)
		out.PackageDetail = nil
		return
	}

	var benefits []dto.PackageBenefit
	for _, a := range raw.Allowances {
		if a.AllowanceType == "PRICE_PLAN" {
			benefits = append(benefits, a.Benefit...)
		}
	}
	out.PackageDetail = benefits
}

func (ctrl *ProfileController) mergeDeviceSpecs(msisdn string, out *dto.MsisdnDetailResponse) {
	res := ctrl.client.Get("/nbss/device-specs-info/")
	if !res.Success {
		logger.Error.Printf("profile: device-specs upstream failed for %s: %v", msisdn, res.Message)
		out.DeviceModel = "-"
		return
	}

	rawBytes, err := json.Marshal(res.Data)
	if err != nil {
		logger.Error.Printf("profile: device-specs marshal failed for %s: %v", msisdn, err)
		out.DeviceModel = "-"
		return
	}

	var raw dto.RawDeviceSpecsData
	if err := json.Unmarshal(rawBytes, &raw); err != nil {
		logger.Error.Printf("profile: device-specs unmarshal failed for %s: %v", msisdn, err)
		out.DeviceModel = "-"
		return
	}

	out.DeviceModel = raw.DeviceModel
}
