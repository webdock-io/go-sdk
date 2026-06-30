package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type DisableIPv6Options struct {
	ServerSlug string
}

type DisableIPv6Response struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) DisableIPv6(ctx context.Context, opts DisableIPv6Options) (*DisableIPv6Response, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/disable-ipv6", opts.ServerSlug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &DisableIPv6Response{CallbackID: callbackID}, nil
}

type EnableIPv6Options struct {
	ServerSlug string
}

type EnableIPv6Response struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) EnableIPv6(ctx context.Context, opts EnableIPv6Options) (*EnableIPv6Response, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/enable-ipv6", opts.ServerSlug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &EnableIPv6Response{CallbackID: callbackID}, nil
}
