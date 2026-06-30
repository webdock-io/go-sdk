package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type WarningDTO struct {
	Type    string         `json:"type" tfsdk:"type"`
	Message string         `json:"message" tfsdk:"message"`
	Data    map[string]any `json:"data" tfsdk:"data"`
}

type ChargeSummaryPriceDTO struct {
	Amount   float64 `json:"amount" tfsdk:"amount"`
	Currency string  `json:"currency" tfsdk:"currency"`
}

type ChargeSummaryItemDTO struct {
	Text     string                `json:"text" tfsdk:"text"`
	Price    ChargeSummaryPriceDTO `json:"price" tfsdk:"price"`
	IsRefund bool                  `json:"isRefund" tfsdk:"is_refund"`
}

type ChargeSummaryTotalDTO struct {
	SubTotal ChargeSummaryPriceDTO `json:"subTotal" tfsdk:"sub_total"`
	VAT      ChargeSummaryPriceDTO `json:"vat" tfsdk:"vat"`
	Total    ChargeSummaryPriceDTO `json:"total" tfsdk:"total"`
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

func (s *Servers) ResizeDryRun(ctx context.Context, opts DryRunResizeServerOptions) (*ResizeDryRunResponse, error) {
	data, err := json.Marshal(map[string]string{"profileSlug": opts.ProfileSlug})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out ResizeDryRunResponse
	_, err = s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/resize/dryrun", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
