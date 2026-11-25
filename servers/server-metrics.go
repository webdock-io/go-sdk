package servers

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
)

type MetricsSamplingDTO struct {
	Amount    float64 `json:"amount" tfsdk:"amount"`
	Timestamp string  `json:"timestamp" tfsdk:"timestamp"`
}

type MetricsPercentSamplingDTO struct {
	Percent   float64 `json:"percent" tfsdk:"percent"`
	Timestamp string  `json:"timestamp" tfsdk:"timestamp"`
}

type MetricsValueSamplingDTO struct {
	Value     float64 `json:"value" tfsdk:"value"`
	Timestamp string  `json:"timestamp" tfsdk:"timestamp"`
}

type DiskDetailedMetricsDTO struct {
	DiskReadBytes        []MetricsValueSamplingDTO `json:"diskReadBytes" tfsdk:"disk_read_bytes"`
	FilesystemSizeBytes  []MetricsValueSamplingDTO `json:"filesystemSizeBytes" tfsdk:"filesystem_size_bytes"`
	FilesystemAvailBytes []MetricsValueSamplingDTO `json:"filesystemAvailBytes" tfsdk:"filesystem_avail_bytes"`
	FilesystemFreeBytes  []MetricsValueSamplingDTO `json:"filesystemFreeBytes" tfsdk:"filesystem_free_bytes"`
	DiskWrittenBytes     []MetricsValueSamplingDTO `json:"diskWrittenBytes" tfsdk:"disk_written_bytes"`
	DiskReadsCompleted   []MetricsValueSamplingDTO `json:"diskReadsCompleted" tfsdk:"disk_reads_completed"`
	DiskWritesCompleted  []MetricsValueSamplingDTO `json:"diskWritesCompleted" tfsdk:"disk_writes_completed"`
}

type NetworkDetailedMetricsDTO struct {
	NetworkReceiveBytes           []MetricsValueSamplingDTO `json:"networkReceiveBytes" tfsdk:"network_receive_bytes"`
	NetworkTransmitBytes          []MetricsValueSamplingDTO `json:"networkTransmitBytes" tfsdk:"network_transmit_bytes"`
	NetworkReceiveDroppedPackets  []MetricsValueSamplingDTO `json:"networkReceiveDroppedPackets" tfsdk:"network_receive_dropped_packets"`
	NetworkTransmitDroppedPackets []MetricsValueSamplingDTO `json:"networkTransmitDroppedPackets" tfsdk:"network_transmit_dropped_packets"`
	NetworkReceiveErrors          []MetricsValueSamplingDTO `json:"networkReceiveErrors" tfsdk:"network_receive_errors"`
	NetworkTransmitErrors         []MetricsValueSamplingDTO `json:"networkTransmitErrors" tfsdk:"network_transmit_errors"`
	NetworkReceivePackets         []MetricsValueSamplingDTO `json:"networkReceivePackets" tfsdk:"network_receive_packets"`
	NetworkTransmitPackets        []MetricsValueSamplingDTO `json:"networkTransmitPackets" tfsdk:"network_transmit_packets"`
}

type CPUDetailedMetricsDTO struct {
	BootTimeSeconds   []MetricsValueSamplingDTO `json:"bootTimeSeconds" tfsdk:"boot_time_seconds"`
	CPUEffectiveTotal []MetricsValueSamplingDTO `json:"cpuEffectiveTotal" tfsdk:"cpu_effective_total"`
}

type MemoryDetailedMetricsDTO struct {
	MemoryActiveBytes         []MetricsValueSamplingDTO `json:"memoryActiveBytes" tfsdk:"memory_active_bytes"`
	MemoryMemAvailableBytes   []MetricsValueSamplingDTO `json:"memoryMemAvailableBytes" tfsdk:"memory_mem_available_bytes"`
	MemoryMemFreeBytes        []MetricsValueSamplingDTO `json:"memoryMemFreeBytes" tfsdk:"memory_mem_free_bytes"`
	MemoryMemTotalBytes       []MetricsValueSamplingDTO `json:"memoryMemTotalBytes" tfsdk:"memory_mem_total_bytes"`
	MemoryActiveFileBytes     []MetricsValueSamplingDTO `json:"memoryActiveFileBytes" tfsdk:"memory_active_file_bytes"`
	MemoryActiveAnonBytes     []MetricsValueSamplingDTO `json:"memoryActiveAnonBytes" tfsdk:"memory_active_anon_bytes"`
	MemoryInactiveBytes       []MetricsValueSamplingDTO `json:"memoryInactiveBytes" tfsdk:"memory_inactive_bytes"`
	MemoryInactiveFileBytes   []MetricsValueSamplingDTO `json:"memoryInactiveFileBytes" tfsdk:"memory_inactive_file_bytes"`
	MemoryInactiveAnonBytes   []MetricsValueSamplingDTO `json:"memoryInactiveAnonBytes" tfsdk:"memory_inactive_anon_bytes"`
	MemoryRssBytes            []MetricsValueSamplingDTO `json:"memoryRssBytes" tfsdk:"memory_rss_bytes"`
	MemoryShmemBytes          []MetricsValueSamplingDTO `json:"memoryShmemBytes" tfsdk:"memory_shmem_bytes"`
	MemorySwapBytes           []MetricsValueSamplingDTO `json:"memorySwapBytes" tfsdk:"memory_swap_bytes"`
	MemoryUnevictableBytes    []MetricsValueSamplingDTO `json:"memoryUnevictableBytes" tfsdk:"memory_unevictable_bytes"`
	MemoryWritebackBytes      []MetricsValueSamplingDTO `json:"memoryWritebackBytes" tfsdk:"memory_writeback_bytes"`
	MemoryCachedBytes         []MetricsValueSamplingDTO `json:"memoryCachedBytes" tfsdk:"memory_cached_bytes"`
	MemoryDirtyBytes          []MetricsValueSamplingDTO `json:"memoryDirtyBytes" tfsdk:"memory_dirty_bytes"`
	MemoryMappedBytes         []MetricsValueSamplingDTO `json:"memoryMappedBytes" tfsdk:"memory_mapped_bytes"`
	MemoryHugepagesTotalBytes []MetricsValueSamplingDTO `json:"memoryHugepagesTotalBytes" tfsdk:"memory_hugepages_total_bytes"`
	MemoryHugepagesFreeBytes  []MetricsValueSamplingDTO `json:"memoryHugepagesFreeBytes" tfsdk:"memory_hugepages_free_bytes"`
	MemoryOomKills            []MetricsValueSamplingDTO `json:"memoryOomKills" tfsdk:"memory_oom_kills"`
}

type DiskMetricsDTO struct {
	Allowed         float64                `json:"allowed" tfsdk:"allowed"`
	Samplings       []MetricsSamplingDTO   `json:"samplings" tfsdk:"samplings"`
	DetailedMetrics DiskDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type NetworkMetricsDTO struct {
	Total            float64                   `json:"total" tfsdk:"total"`
	Allowed          float64                   `json:"allowed" tfsdk:"allowed"`
	IngressSamplings []MetricsSamplingDTO      `json:"ingressSamplings" tfsdk:"ingress_samplings"`
	EgressSamplings  []MetricsSamplingDTO      `json:"egressSamplings" tfsdk:"egress_samplings"`
	DetailedMetrics  NetworkDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type CPUMetricsDTO struct {
	UsageSamplings        []MetricsSamplingDTO        `json:"usageSamplings" tfsdk:"usage_samplings"`
	UsagePercentSamplings []MetricsPercentSamplingDTO `json:"usagePercentSamplings" tfsdk:"usage_percent_samplings"`
	DetailedMetrics       CPUDetailedMetricsDTO       `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type ProcessesMetricsDTO struct {
	ProcessesSamplings []MetricsSamplingDTO `json:"processesSamplings" tfsdk:"processes_samplings"`
}

type MemoryMetricsDTO struct {
	UsageSamplings  []MetricsSamplingDTO     `json:"usageSamplings" tfsdk:"usage_samplings"`
	DetailedMetrics MemoryDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type ResourceUsageLevel string

const (
	ResourceUsageOK   ResourceUsageLevel = "ok"
	ResourceUsageWarn ResourceUsageLevel = "warn"
	ResourceUsageCrit ResourceUsageLevel = "crit"
)

type ResourceUsageMetricStatusDTO struct {
	Used                     float64            `json:"used" tfsdk:"used"`
	Allowed                  float64            `json:"allowed" tfsdk:"allowed"`
	PercentUsed              float64            `json:"percentUsed" tfsdk:"percent_used"`
	WarnThresholdPercent     float64            `json:"warnThresholdPercent" tfsdk:"warn_threshold_percent"`
	CriticalThresholdPercent float64            `json:"criticalThresholdPercent" tfsdk:"critical_threshold_percent"`
	ThresholdHit             bool               `json:"thresholdHit" tfsdk:"threshold_hit"`
	Level                    ResourceUsageLevel `json:"level" tfsdk:"level"`
}

type ResourceUsageProblemResourceDTO struct {
	// Resource is one of network, disk, memory, cpu, memoryOomKills,
	// networkDroppedPackets or networkErrors.
	Resource string             `json:"resource" tfsdk:"resource"`
	Level    ResourceUsageLevel `json:"level" tfsdk:"level"`
	Evidence map[string]any     `json:"evidence" tfsdk:"evidence"`
}

type ResourceUsageStatusDTO struct {
	HasIssues        bool                              `json:"hasIssues" tfsdk:"has_issues"`
	Level            ResourceUsageLevel                `json:"level" tfsdk:"level"`
	Network          ResourceUsageMetricStatusDTO      `json:"network" tfsdk:"network"`
	Disk             ResourceUsageMetricStatusDTO      `json:"disk" tfsdk:"disk"`
	Memory           ResourceUsageMetricStatusDTO      `json:"memory" tfsdk:"memory"`
	CPU              ResourceUsageMetricStatusDTO      `json:"cpu" tfsdk:"cpu"`
	ProblemResources []ResourceUsageProblemResourceDTO `json:"problemResources" tfsdk:"problem_resources"`
}

type ServerMetricsDTO struct {
	Disk                DiskMetricsDTO         `json:"disk" tfsdk:"disk"`
	Network             NetworkMetricsDTO      `json:"network" tfsdk:"network"`
	CPU                 CPUMetricsDTO          `json:"cpu" tfsdk:"cpu"`
	Processes           ProcessesMetricsDTO    `json:"processes" tfsdk:"processes"`
	Memory              MemoryMetricsDTO       `json:"memory" tfsdk:"memory"`
	ResourceUsageStatus ResourceUsageStatusDTO `json:"resourceUsageStatus" tfsdk:"resource_usage_status"`
	Meta                map[string]any         `json:"meta" tfsdk:"meta"`
}

// Instant* detailed metrics hold the latest single sampling per metric.

type InstantDiskDetailedMetricsDTO struct {
	DiskReadBytes        MetricsValueSamplingDTO `json:"diskReadBytes" tfsdk:"disk_read_bytes"`
	FilesystemSizeBytes  MetricsValueSamplingDTO `json:"filesystemSizeBytes" tfsdk:"filesystem_size_bytes"`
	FilesystemAvailBytes MetricsValueSamplingDTO `json:"filesystemAvailBytes" tfsdk:"filesystem_avail_bytes"`
	FilesystemFreeBytes  MetricsValueSamplingDTO `json:"filesystemFreeBytes" tfsdk:"filesystem_free_bytes"`
	DiskWrittenBytes     MetricsValueSamplingDTO `json:"diskWrittenBytes" tfsdk:"disk_written_bytes"`
	DiskReadsCompleted   MetricsValueSamplingDTO `json:"diskReadsCompleted" tfsdk:"disk_reads_completed"`
	DiskWritesCompleted  MetricsValueSamplingDTO `json:"diskWritesCompleted" tfsdk:"disk_writes_completed"`
}

type InstantNetworkDetailedMetricsDTO struct {
	NetworkReceiveBytes           MetricsValueSamplingDTO `json:"networkReceiveBytes" tfsdk:"network_receive_bytes"`
	NetworkTransmitBytes          MetricsValueSamplingDTO `json:"networkTransmitBytes" tfsdk:"network_transmit_bytes"`
	NetworkReceiveDroppedPackets  MetricsValueSamplingDTO `json:"networkReceiveDroppedPackets" tfsdk:"network_receive_dropped_packets"`
	NetworkTransmitDroppedPackets MetricsValueSamplingDTO `json:"networkTransmitDroppedPackets" tfsdk:"network_transmit_dropped_packets"`
	NetworkReceiveErrors          MetricsValueSamplingDTO `json:"networkReceiveErrors" tfsdk:"network_receive_errors"`
	NetworkTransmitErrors         MetricsValueSamplingDTO `json:"networkTransmitErrors" tfsdk:"network_transmit_errors"`
	NetworkReceivePackets         MetricsValueSamplingDTO `json:"networkReceivePackets" tfsdk:"network_receive_packets"`
	NetworkTransmitPackets        MetricsValueSamplingDTO `json:"networkTransmitPackets" tfsdk:"network_transmit_packets"`
}

type InstantCPUDetailedMetricsDTO struct {
	BootTimeSeconds   MetricsValueSamplingDTO `json:"bootTimeSeconds" tfsdk:"boot_time_seconds"`
	CPUEffectiveTotal MetricsValueSamplingDTO `json:"cpuEffectiveTotal" tfsdk:"cpu_effective_total"`
}

type InstantMemoryDetailedMetricsDTO struct {
	MemoryActiveBytes         MetricsValueSamplingDTO `json:"memoryActiveBytes" tfsdk:"memory_active_bytes"`
	MemoryMemAvailableBytes   MetricsValueSamplingDTO `json:"memoryMemAvailableBytes" tfsdk:"memory_mem_available_bytes"`
	MemoryMemFreeBytes        MetricsValueSamplingDTO `json:"memoryMemFreeBytes" tfsdk:"memory_mem_free_bytes"`
	MemoryMemTotalBytes       MetricsValueSamplingDTO `json:"memoryMemTotalBytes" tfsdk:"memory_mem_total_bytes"`
	MemoryActiveFileBytes     MetricsValueSamplingDTO `json:"memoryActiveFileBytes" tfsdk:"memory_active_file_bytes"`
	MemoryActiveAnonBytes     MetricsValueSamplingDTO `json:"memoryActiveAnonBytes" tfsdk:"memory_active_anon_bytes"`
	MemoryInactiveBytes       MetricsValueSamplingDTO `json:"memoryInactiveBytes" tfsdk:"memory_inactive_bytes"`
	MemoryInactiveFileBytes   MetricsValueSamplingDTO `json:"memoryInactiveFileBytes" tfsdk:"memory_inactive_file_bytes"`
	MemoryInactiveAnonBytes   MetricsValueSamplingDTO `json:"memoryInactiveAnonBytes" tfsdk:"memory_inactive_anon_bytes"`
	MemoryRssBytes            MetricsValueSamplingDTO `json:"memoryRssBytes" tfsdk:"memory_rss_bytes"`
	MemoryShmemBytes          MetricsValueSamplingDTO `json:"memoryShmemBytes" tfsdk:"memory_shmem_bytes"`
	MemorySwapBytes           MetricsValueSamplingDTO `json:"memorySwapBytes" tfsdk:"memory_swap_bytes"`
	MemoryUnevictableBytes    MetricsValueSamplingDTO `json:"memoryUnevictableBytes" tfsdk:"memory_unevictable_bytes"`
	MemoryWritebackBytes      MetricsValueSamplingDTO `json:"memoryWritebackBytes" tfsdk:"memory_writeback_bytes"`
	MemoryCachedBytes         MetricsValueSamplingDTO `json:"memoryCachedBytes" tfsdk:"memory_cached_bytes"`
	MemoryDirtyBytes          MetricsValueSamplingDTO `json:"memoryDirtyBytes" tfsdk:"memory_dirty_bytes"`
	MemoryMappedBytes         MetricsValueSamplingDTO `json:"memoryMappedBytes" tfsdk:"memory_mapped_bytes"`
	MemoryHugepagesTotalBytes MetricsValueSamplingDTO `json:"memoryHugepagesTotalBytes" tfsdk:"memory_hugepages_total_bytes"`
	MemoryHugepagesFreeBytes  MetricsValueSamplingDTO `json:"memoryHugepagesFreeBytes" tfsdk:"memory_hugepages_free_bytes"`
	MemoryOomKills            MetricsValueSamplingDTO `json:"memoryOomKills" tfsdk:"memory_oom_kills"`
}

type InstantDiskMetricsDTO struct {
	Allowed         float64                       `json:"allowed" tfsdk:"allowed"`
	LastSamplings   MetricsSamplingDTO            `json:"lastSamplings" tfsdk:"last_samplings"`
	DetailedMetrics InstantDiskDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type InstantNetworkMetricsDTO struct {
	Total                 float64                          `json:"total" tfsdk:"total"`
	Allowed               float64                          `json:"allowed" tfsdk:"allowed"`
	LatestIngressSampling MetricsSamplingDTO               `json:"latestIngressSampling" tfsdk:"latest_ingress_sampling"`
	LatestEgressSampling  MetricsSamplingDTO               `json:"latestEgressSampling" tfsdk:"latest_egress_sampling"`
	DetailedMetrics       InstantNetworkDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type InstantCPUMetricsDTO struct {
	LatestUsageSampling        MetricsSamplingDTO           `json:"latestUsageSampling" tfsdk:"latest_usage_sampling"`
	LatestUsagePercentSampling MetricsPercentSamplingDTO    `json:"latestUsagePercentSampling" tfsdk:"latest_usage_percent_sampling"`
	DetailedMetrics            InstantCPUDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type InstantProcessesMetricsDTO struct {
	LatestProcessesSampling MetricsSamplingDTO `json:"latestProcessesSampling" tfsdk:"latest_processes_sampling"`
}

type InstantMemoryMetricsDTO struct {
	LatestUsageSampling MetricsSamplingDTO              `json:"latestUsageSampling" tfsdk:"latest_usage_sampling"`
	DetailedMetrics     InstantMemoryDetailedMetricsDTO `json:"detailedMetrics" tfsdk:"detailed_metrics"`
}

type InstantServerMetricsDTO struct {
	Disk                InstantDiskMetricsDTO      `json:"disk" tfsdk:"disk"`
	Network             InstantNetworkMetricsDTO   `json:"network" tfsdk:"network"`
	CPU                 InstantCPUMetricsDTO       `json:"cpu" tfsdk:"cpu"`
	Processes           InstantProcessesMetricsDTO `json:"processes" tfsdk:"processes"`
	Memory              InstantMemoryMetricsDTO    `json:"memory" tfsdk:"memory"`
	ResourceUsageStatus ResourceUsageStatusDTO     `json:"resourceUsageStatus" tfsdk:"resource_usage_status"`
}

type ServerMetricsWarningServerDTO struct {
	Slug         string `json:"slug" tfsdk:"slug"`
	Name         string `json:"name" tfsdk:"name"`
	LocationName string `json:"locationName" tfsdk:"location_name"`
	LocationCity string `json:"locationCity" tfsdk:"location_city"`
}

type ServerMetricsWarningRowDTO struct {
	Server              ServerMetricsWarningServerDTO `json:"server" tfsdk:"server"`
	ResourceUsageStatus ResourceUsageStatusDTO        `json:"resourceUsageStatus" tfsdk:"resource_usage_status"`
}

type ServerMetricsWarningsOverviewDTO struct {
	Count int                          `json:"count" tfsdk:"count"`
	Rows  []ServerMetricsWarningRowDTO `json:"rows" tfsdk:"rows"`
}

type GetServerMetricsOptions struct {
	ServerSlug        string
	ResolutionSeconds *int
	MaxSamples        *int
	// From and To accept an ISO 8601 timestamp or a Unix timestamp as a string.
	From string
	To   string
}

// GetServerMetrics mirrors servers.getServerMetrics() in the Node SDK.
func (s *Servers) GetServerMetrics(ctx context.Context, opts GetServerMetricsOptions) (*ServerMetricsDTO, error) {
	q := url.Values{}
	if opts.ResolutionSeconds != nil {
		q.Set("resolutionSeconds", strconv.Itoa(*opts.ResolutionSeconds))
	}
	if opts.MaxSamples != nil {
		q.Set("maxSamples", strconv.Itoa(*opts.MaxSamples))
	}
	if opts.From != "" {
		q.Set("from", opts.From)
	}
	if opts.To != "" {
		q.Set("to", opts.To)
	}
	u := &url.URL{Path: fmt.Sprintf("/servers/%s/metrics", opts.ServerSlug), RawQuery: q.Encode()}

	var out ServerMetricsDTO
	if _, err := s.client.DoJSON(ctx, "GET", u.String(), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type GetInstantMetricsOptions struct {
	ServerSlug string
}

// GetInstantMetrics mirrors servers.getInstantMetrics() in the Node SDK.
func (s *Servers) GetInstantMetrics(ctx context.Context, opts GetInstantMetricsOptions) (*InstantServerMetricsDTO, error) {
	var out InstantServerMetricsDTO
	if _, err := s.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/metrics/now", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type GetServerMetricsWarningsOverviewOptions struct {
	// ServerSlug limits the overview to one server. Empty means all servers.
	ServerSlug string
	IncludeOk  *bool
}

// GetServerMetricsWarningsOverview mirrors
// servers.getServerMetricsWarningsOverview() in the Node SDK.
func (s *Servers) GetServerMetricsWarningsOverview(ctx context.Context, opts ...GetServerMetricsWarningsOverviewOptions) (*ServerMetricsWarningsOverviewDTO, error) {
	var options GetServerMetricsWarningsOverviewOptions
	if len(opts) > 0 {
		options = opts[0]
	}
	q := url.Values{}
	if options.ServerSlug != "" {
		q.Set("serverSlug", options.ServerSlug)
	}
	if options.IncludeOk != nil {
		q.Set("includeOk", strconv.FormatBool(*options.IncludeOk))
	}
	u := &url.URL{Path: "/servers/metrics/warnings", RawQuery: q.Encode()}

	var out ServerMetricsWarningsOverviewDTO
	if _, err := s.client.DoJSON(ctx, "GET", u.String(), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
