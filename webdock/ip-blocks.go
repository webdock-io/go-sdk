package platform

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/webdock-io/go-sdk/client"
)

type WebdockIPBlocks struct {
	client *client.Client
}

func NewWebdockIPBlocks(c *client.Client) WebdockIPBlocks {
	return WebdockIPBlocks{client: c}
}

type IPBlockStatus string

const (
	IPBlockStatusAll      IPBlockStatus = "all"
	IPBlockStatusFree     IPBlockStatus = "free"
	IPBlockStatusUsed     IPBlockStatus = "used"
	IPBlockStatusBanned   IPBlockStatus = "banned"
	IPBlockStatusReserved IPBlockStatus = "reserved"
)

type WebdockIPBlock struct {
	ID         int64         `json:"id" tfsdk:"id"`
	BlockID    int64         `json:"blockId" tfsdk:"block_id"`
	IPv4       string        `json:"ipv4" tfsdk:"ipv4"`
	IPv6       string        `json:"ipv6" tfsdk:"ipv6"`
	Status     IPBlockStatus `json:"status" tfsdk:"status"`
	ServerSlug string        `json:"serverSlug" tfsdk:"server_slug"`
}

type BanIPOptions struct {
	IPID int64
}

func (i *WebdockIPBlocks) BanIP(ctx context.Context, opts BanIPOptions) (*WebdockIPBlock, error) {
	return i.setBanned(ctx, opts.IPID, true)
}

func (i *WebdockIPBlocks) UnbanIP(ctx context.Context, opts BanIPOptions) (*WebdockIPBlock, error) {
	return i.setBanned(ctx, opts.IPID, false)
}

func (i *WebdockIPBlocks) setBanned(ctx context.Context, ipID int64, banned bool) (*WebdockIPBlock, error) {
	body, err := json.Marshal(map[string]bool{"banned": banned})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	var out WebdockIPBlock
	_, err = i.client.Do(ctx, "PATCH", fmt.Sprintf("/servers/ipBlocks/%d/banned", ipID), bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type InspectIPBlockOptions struct {
	BlockID int64
	Status  IPBlockStatus
}

func (i *WebdockIPBlocks) Inspect(ctx context.Context, opts InspectIPBlockOptions) ([]WebdockIPBlock, error) {
	status := opts.Status
	if status == "" {
		status = IPBlockStatusAll
	}

	u := &url.URL{Path: "/servers/ipBlocks/inspect"}
	q := url.Values{}
	q.Set("blockId", fmt.Sprintf("%d", opts.BlockID))
	q.Set("status", string(status))
	u.RawQuery = q.Encode()

	var raw json.RawMessage
	_, err := i.client.Do(ctx, "GET", u.String(), nil, &raw)
	if err != nil {
		return nil, err
	}

	var list []WebdockIPBlock
	if err := json.Unmarshal(raw, &list); err == nil {
		return list, nil
	}

	var single WebdockIPBlock
	if err := json.Unmarshal(raw, &single); err != nil {
		return nil, err
	}
	return []WebdockIPBlock{single}, nil
}

func (i *WebdockIPBlocks) InspectIPBlock(ctx context.Context, opts InspectIPBlockOptions) ([]WebdockIPBlock, error) {
	return i.Inspect(ctx, opts)
}
