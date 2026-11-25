package servers

import (
	"context"
	"fmt"
)

type StopServerOptions struct {
	Slug string
}

type StopServerResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Stop mirrors servers.stop() in the Node SDK.
func (s *Servers) Stop(ctx context.Context, opts StopServerOptions) (*StopServerResponse, error) {
	var out Server
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/actions/stop", opts.Slug), nil, &out)
	if err != nil {
		return nil, err
	}
	return &StopServerResponse{Server: out, CallbackID: callbackID}, nil
}
