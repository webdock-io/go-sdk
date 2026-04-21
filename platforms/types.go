package platforms

import "github.com/webdock-io/go-sdk/client"

type ResourceLimit struct {
	Min       int `json:"min" tfsdk:"min"`
	Max       int `json:"max" tfsdk:"max"`
	CostCents int `json:"costCents" tfsdk:"cost_cents"`
	FreeUnits int `json:"freeUnits" tfsdk:"free_units"`
}

type PlatformResourceLimits struct {
	CPUThreads       ResourceLimit `json:"cpuThreads" tfsdk:"cpu_threads"`
	RAM              ResourceLimit `json:"ram" tfsdk:"ram"`
	DiskSpace        ResourceLimit `json:"diskSpace" tfsdk:"disk_space"`
	NetworkBandwidth ResourceLimit `json:"networkBandwidth" tfsdk:"network_bandwidth"`
}

type Platform struct {
	Slug           string                 `json:"slug" tfsdk:"slug"`
	Name           string                 `json:"name" tfsdk:"name"`
	Description    *string                `json:"description" tfsdk:"description"`
	ResourceLimits PlatformResourceLimits `json:"resourceLimits" tfsdk:"resource_limits"`
}

type Platforms struct {
	client *client.Client
}

func New(c *client.Client) Platforms {
	return Platforms{client: c}
}
