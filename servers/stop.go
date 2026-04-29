package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type StopServerOptions struct {
	Slug string
}

type StopServerResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Stop(ctx context.Context, opts StopServerOptions) (*StopServerResponse, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/stop", opts.Slug), nil, nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &StopServerResponse{CallbackID: callbackID}, nil
}
