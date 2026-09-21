package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/your-org/ums-bff-service-customer-360/pkg/logger"
	"github.com/your-org/ums-bff-service-customer-360/pkg/response"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type thirdPartyResult struct {
	ResultCode        string `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
}

type thirdPartyResponse struct {
	Result thirdPartyResult `json:"result"`
	Data   json.RawMessage  `json:"responseData"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) Get(path string) response.Response {
	start := time.Now()
	url := c.baseURL + path

	logger.Info.Printf("[THIRD-PARTY] GET %s", url)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		logger.Error.Printf("[THIRD-PARTY] GET %s error: %v (duration: %s)", url, err, time.Since(start))
		return response.Error(fmt.Sprintf("failed to call third-party: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Printf("[THIRD-PARTY] GET %s read body error: %v (duration: %s)", url, err, time.Since(start))
		return response.Error(fmt.Sprintf("failed to read response: %v", err))
	}

	logger.Info.Printf("[THIRD-PARTY] GET %s status=%d duration=%s body=%s", url, resp.StatusCode, time.Since(start), string(body))

	if resp.StatusCode != http.StatusOK {
		return response.Error(fmt.Sprintf("third-party returned status %d", resp.StatusCode))
	}

	var tpResp thirdPartyResponse
	if err := json.Unmarshal(body, &tpResp); err != nil {
		logger.Error.Printf("[THIRD-PARTY] GET %s unmarshal error: %v", url, err)
		return response.Error(fmt.Sprintf("failed to parse response: %v", err))
	}

	if tpResp.Result.ResultCode != "00" {
		logger.Info.Printf("[THIRD-PARTY] GET %s resultCode=%s description=%s", url, tpResp.Result.ResultCode, tpResp.Result.ResultDescription)
		return response.Error(tpResp.Result.ResultDescription)
	}

	var data interface{}
	if err := json.Unmarshal(tpResp.Data, &data); err != nil {
		logger.Error.Printf("[THIRD-PARTY] GET %s unmarshal responseData error: %v", url, err)
		return response.Error(fmt.Sprintf("failed to parse responseData: %v", err))
	}

	return response.Success(data, "success")
}
