package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CancelDeleteOptions struct {
	ServerSlug string
}

type CancelDeleteResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) CancelDelete(ctx context.Context, opts CancelDeleteOptions) (*CancelDeleteResponse, error) {
	var out Server
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/uncancel", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CancelDeleteResponse{Server: out, CallbackID: callbackID}, nil
}
