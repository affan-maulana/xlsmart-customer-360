package dto

type GetDeviceSpecsResponse struct {
	Manufacturer string `json:"manufacturer"`
	ModelName    string `json:"modelName"`
	OsName       string `json:"osName"`
	OsVersion    string `json:"osVersion"`
	Imei         string `json:"imei"`
}

type RawDeviceSpecsData struct {
	BEARER_EDGE     string `json:"BEARER_EDGE"`
	BEARER_GSM_CSD  string `json:"BEARER_GSM_CSD"`
	BEARER_GSM_GPRS string `json:"BEARER_GSM_GPRS"`
	BEARER_GSM_SMS  string `json:"BEARER_GSM_SMS"`
	BEARER_HSDPA    string `json:"BEARER_HSDPA"`
	BEARER_HSUPA    string `json:"BEARER_HSUPA"`
	BEARER_LTE      string `json:"BEARER_LTE"`
	BEARER_UMTS     string `json:"BEARER_UMTS"`
	BEARER_WLAN     string `json:"BEARER_WLAN"`
	BROWSER_NAME    string `json:"BROWSER_NAME"`
	FDD_BAND_3      string `json:"FDD_BAND_3"`
	FDD_BAND_8      string `json:"FDD_BAND_8"`
	HARDWARE_TYPE   string `json:"HARDWARE_TYPE"`
	IMEI            string `json:"IMEI"`
	IMSI            string `json:"IMSI"`
	OS_NAME         string `json:"OS_NAME"`
	OS_VERSION      string `json:"OS_VERSION"`
	TCR_MODEL       string `json:"TCR_MODEL"`
	TCR_NAME        string `json:"TCR_NAME"`
}

type RawDeviceSpecsResponse struct {
	Code    string             `json:"code"`
	Data    RawDeviceSpecsData `json:"data"`
	Message string             `json:"message"`
	Status  string             `json:"status"`
}
