package gosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WarningDTO struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ChargeSummaryItemDTO struct {
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Currency    string  `json:"currency,omitempty"`
}

type ChargeSummaryTotalDTO struct {
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type ChargeSummaryDTO struct {
	Items    []ChargeSummaryItemDTO `json:"items"`
	IsRefund bool                   `json:"isRefund"`
	Total    ChargeSummaryTotalDTO  `json:"total"`
}

type ResizeDryRunResponse struct {
	Warnings      []WarningDTO     `json:"warnings"`
	ChargeSummary ChargeSummaryDTO `json:"chargeSummary"`
}

type ResizeDryRunRequest struct {
	ProfileSlug string `json:"profileSlug"`
}

func (w *Webdock) DryRunResizeServer(serverSlug string, profileSlug string) (ResizeDryRunResponse, error) {
	URL := url.URL{
		Scheme: "https",
		Host:   w.BASE_URL,
		Path:   fmt.Sprintf(`/v1/servers/%s/actions/resize/dryrun`, serverSlug),
	}

	requestBody := ResizeDryRunRequest{
		ProfileSlug: profileSlug,
	}

	data, err := json.Marshal(requestBody)
	if err != nil {
		return ResizeDryRunResponse{}, fmt.Errorf("error marshaling request body: %w", err)
	}

	req, err := http.NewRequest("POST", URL.String(), bytes.NewBuffer(data))
	req.Header.Set(w.Authorization, w.GetFormatedToken())
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return ResizeDryRunResponse{}, fmt.Errorf("error creating request: %w", err)
	}

	res, err := w.client.Do(req)
	if err != nil {
		return ResizeDryRunResponse{}, fmt.Errorf("operation failed: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResizeDryRunResponse{}, fmt.Errorf("failed to read response: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		webdock := WebdockError{
			ID:      1,
			Message: "error occurred",
		}
		err = json.Unmarshal(body, &webdock)
		if err != nil {
			return ResizeDryRunResponse{}, fmt.Errorf("%s", http.StatusText(res.StatusCode))
		}
		return ResizeDryRunResponse{}, fmt.Errorf("operation failed: %s", webdock.Message)
	}

	var response ResizeDryRunResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return ResizeDryRunResponse{}, fmt.Errorf("error decoding response: %w", err)
	}

	return response, nil
}
