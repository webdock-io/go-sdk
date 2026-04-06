package servers

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type WarningDTO struct {
	Type    string      `json:"type" tfsdk:"type"`
	Message string      `json:"message" tfsdk:"message"`
	Data    interface{} `json:"data" tfsdk:"data"`
}

type ChargeSummaryItemDTO struct {
	Description string  `json:"description,omitempty" tfsdk:"description"`
	Amount      float64 `json:"amount,omitempty" tfsdk:"amount"`
	Currency    string  `json:"currency,omitempty" tfsdk:"currency"`
}

type ChargeSummaryTotalDTO struct {
	Amount   float64 `json:"amount,omitempty" tfsdk:"amount"`
	Currency string  `json:"currency,omitempty" tfsdk:"currency"`
}

type ChargeSummaryDTO struct {
	Items    []ChargeSummaryItemDTO `json:"items" tfsdk:"items"`
	IsRefund bool                   `json:"isRefund" tfsdk:"is_refund"`
	Total    ChargeSummaryTotalDTO  `json:"total" tfsdk:"total"`
}

type ResizeDryRunResponse struct {
	Warnings      []WarningDTO     `json:"warnings" tfsdk:"warnings"`
	ChargeSummary ChargeSummaryDTO `json:"chargeSummary" tfsdk:"charge_summary"`
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
