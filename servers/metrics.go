package servers

import (
	"context"
	"fmt"
)

type MetricSampling struct {
	Amount    float64 `json:"amount" tfsdk:"amount"`
	Timestamp string  `json:"timestamp" tfsdk:"timestamp"`
}

type DiskMetrics struct {
	Allowed   float64          `json:"allowed" tfsdk:"allowed"`
	Samplings []MetricSampling `json:"samplings" tfsdk:"samplings"`
}

type NetworkMetrics struct {
	Total            float64          `json:"total" tfsdk:"total"`
	Allowed          float64          `json:"allowed" tfsdk:"allowed"`
	IngressSamplings []MetricSampling `json:"ingressSamplings" tfsdk:"ingress_samplings"`
	EgressSamplings  []MetricSampling `json:"egressSamplings" tfsdk:"egress_samplings"`
}

type CPUMetrics struct {
	UsageSamplings []MetricSampling `json:"usageSamplings" tfsdk:"usage_samplings"`
}

type ProcessesMetrics struct {
	ProcessesSamplings []MetricSampling `json:"processesSamplings" tfsdk:"processes_samplings"`
}

type MemoryMetrics struct {
	UsageSamplings []MetricSampling `json:"usageSamplings" tfsdk:"usage_samplings"`
}

type MetricsResponse struct {
	Disk      DiskMetrics      `json:"disk" tfsdk:"disk"`
	Network   NetworkMetrics   `json:"network" tfsdk:"network"`
	CPU       CPUMetrics       `json:"cpu" tfsdk:"cpu"`
	Processes ProcessesMetrics `json:"processes" tfsdk:"processes"`
	Memory    MemoryMetrics    `json:"memory" tfsdk:"memory"`
}

type MetricsOptions struct {
	ServerSlug string
	Now        bool
}

func (s *Servers) Metrics(ctx context.Context, opts MetricsOptions) (*MetricsResponse, error) {
	endpoint := fmt.Sprintf("/servers/%s/metrics", opts.ServerSlug)
	if opts.Now {
		endpoint += "/now"
	}

	var out MetricsResponse
	_, err := s.client.Do(ctx, "GET", endpoint, nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
