package servers

import (
	"bytes"
	"context"
	"encoding/json"
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

func (s *ServerIPBlocks) ListIPBlocks(ctx context.Context, opts ListIPBlocksOptions) (map[int64]IPBlockStats, error) {
	return s.List(ctx, opts)
}

type ChangeIPAddressOptions struct {
	ServerSlug           string `json:"-" tfsdk:"server_slug"`
	SameBlock            bool   `json:"sameBlock" tfsdk:"same_block"`
	OverrideBlockID      int64  `json:"overrideBlockId" tfsdk:"override_block_id"`
	MarkReleasedIPBanned bool   `json:"markReleasedIpBanned" tfsdk:"mark_released_ip_banned"`
}

type ChangeIPAddressResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *ServerIPBlocks) ChangeIPAddress(ctx context.Context, opts ChangeIPAddressOptions) (*ChangeIPAddressResponse, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/changeIp", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &ChangeIPAddressResponse{CallbackID: callbackID}, nil
}

func (s *ServerIPBlocks) ChangeIP(ctx context.Context, opts ChangeIPAddressOptions) (*ChangeIPAddressResponse, error) {
	return s.ChangeIPAddress(ctx, opts)
}
