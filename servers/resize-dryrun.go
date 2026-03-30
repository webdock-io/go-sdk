package servers

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type DryRunResizeServerOptions struct {
	ServerSlug  string
	ProfileSlug string
}

func (s *Servers) ResizeDryRun(opts DryRunResizeServerOptions) (*ResizeDryRunResponse, error) {
	data, err := json.Marshal(map[string]string{"profileSlug": opts.ProfileSlug})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ResizeDryRunResponse
	_, err = s.client.Do("POST", fmt.Sprintf("v1/servers/%s/actions/resize/dryrun", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
