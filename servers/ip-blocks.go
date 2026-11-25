package servers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type ServerIPBlocks struct {
	client *client.Client
}

func NewServerIPBlocks(c *client.Client) ServerIPBlocks {
	return ServerIPBlocks{client: c}
}

type IPBlockStats struct {
	CIDR     string `json:"cidr" tfsdk:"cidr"`
	Total    int    `json:"total" tfsdk:"total"`
	Free     int    `json:"free" tfsdk:"free"`
	Used     int    `json:"used" tfsdk:"used"`
	Banned   int    `json:"banned" tfsdk:"banned"`
	Reserved int    `json:"reserved" tfsdk:"reserved"`
}

type ListIPBlocksOptions struct {
	ServerSlug string
}

func (s *ServerIPBlocks) List(ctx context.Context, opts ListIPBlocksOptions) (map[int64]IPBlockStats, error) {
	var raw map[string]IPBlockStats
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("/servers/%s/ipBlocks", opts.ServerSlug), nil, &raw)
	if err != nil {
		return nil, err
	}

	out := make(map[int64]IPBlockStats, len(raw))
	for key, stats := range raw {
		blockID, err := strconv.ParseInt(key, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("decoding block id %q: %w", key, err)
		}
		out[blockID] = stats
	}
	return out, nil
}

// ListIPBlocks mirrors servers.IpBlock.ListIpBlocks() in the Node SDK.
func (s *ServerIPBlocks) ListIPBlocks(ctx context.Context, opts ListIPBlocksOptions) (map[int64]IPBlockStats, error) {
	return s.List(ctx, opts)
}

// ServerIP is one IP address assigned to a server.
type ServerIP struct {
	// ID is the IP row ID, as used by ChangeIPAddress and Webdock.IPBlocks.InspectIPBlock.
	ID int64 `json:"id" tfsdk:"id"`
	// BlockID is the IP block this address belongs to (see ListIPBlocks).
	BlockID int64 `json:"blockId" tfsdk:"block_id"`
	// IPv4 is nil for IPv6-only assignments.
	IPv4 *string `json:"ipv4" tfsdk:"ipv4"`
	// IPv6 is nil when the block has no IPv6.
	IPv6 *string `json:"ipv6" tfsdk:"ipv6"`
	// Primary is true for the server's primary (outbound) IP. Exactly one IP per server is primary.
	Primary bool `json:"primary" tfsdk:"primary"`
	// Billable is true when this IP is billed in addition to the server price.
	Billable bool `json:"billable" tfsdk:"billable"`
	// ReverseDNS is the custom PTR hostname, or nil for the default.
	ReverseDNS *string `json:"reverseDns" tfsdk:"reverse_dns"`
}

type ListIPsOptions struct {
	ServerSlug string
}

// ListIPs lists the IPs assigned to a server. Mirrors servers.IpBlock.listIps() in the Node SDK.
func (s *ServerIPBlocks) ListIPs(ctx context.Context, opts ListIPsOptions) ([]ServerIP, error) {
	var out []ServerIP
	if _, err := s.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/ips", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ServerIPToAdd describes an IP to attach to a server. Set exactly one of
// IPID (a specific free IP) or BlockID (any free IP from that block).
type ServerIPToAdd struct {
	IPID    *int64 `json:"ipId,omitempty" tfsdk:"ip_id"`
	BlockID *int64 `json:"blockId,omitempty" tfsdk:"block_id"`
	Primary *bool  `json:"primary,omitempty" tfsdk:"primary"`
	// ReverseDNS sets the PTR hostname; client.Null[string]() keeps the default.
	ReverseDNS client.Nullable[string] `json:"reverseDns,omitzero" tfsdk:"-"`
}

// ReverseDNSUpdate sets the PTR hostname of an IP. A nil Hostname resets it
// to the default.
type ReverseDNSUpdate struct {
	IPID     int64   `json:"ipId" tfsdk:"ip_id"`
	Hostname *string `json:"hostname" tfsdk:"hostname"`
}

type ChangeIPAddressOptions struct {
	ServerSlug      string                 `json:"-" tfsdk:"server_slug"`
	Add             []ServerIPToAdd        `json:"add,omitzero" tfsdk:"add"`
	Remove          []int64                `json:"remove,omitzero" tfsdk:"remove"`
	BanRemovedIPIDs []int64                `json:"banRemovedIpIds,omitzero" tfsdk:"ban_removed_ip_ids"`
	PrimaryIPID     client.Nullable[int64] `json:"primaryIpId,omitzero" tfsdk:"-"`
	ReverseDNS      []ReverseDNSUpdate     `json:"reverseDns,omitzero" tfsdk:"reverse_dns"`
}

type ChangeIPAddressResponse struct {
	IPs        []ServerIP `json:"ips" tfsdk:"ips"`
	CallbackID string     `json:"callbackId" tfsdk:"callback_id"`
}

// ChangeIPAddress adds, removes, re-prioritises and sets reverse DNS for a
// server's IPs. Mirrors servers.IpBlock.ChangeIPAddress() in the Node SDK.
func (s *ServerIPBlocks) ChangeIPAddress(ctx context.Context, opts ChangeIPAddressOptions) (*ChangeIPAddressResponse, error) {
	var out []ServerIP
	callbackID, err := s.client.DoJSON(ctx, "PATCH", fmt.Sprintf("/servers/%s/ips", opts.ServerSlug), opts, &out)
	if err != nil {
		return nil, err
	}
	return &ChangeIPAddressResponse{IPs: out, CallbackID: callbackID}, nil
}

// Deprecated: use ChangeIPAddress.
func (s *ServerIPBlocks) ChangeIP(ctx context.Context, opts ChangeIPAddressOptions) (*ChangeIPAddressResponse, error) {
	return s.ChangeIPAddress(ctx, opts)
}
