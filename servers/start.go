package servers

import (
	"context"
	"fmt"
)

type StartServerOptions struct {
	Slug string
}

type StartServerResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Start mirrors servers.start() in the Node SDK.
func (s *Servers) Start(ctx context.Context, opts StartServerOptions) (*StartServerResponse, error) {
	var out Server
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/actions/start", opts.Slug), nil, &out)
	if err != nil {
		return nil, err
	}
	return &StartServerResponse{Server: out, CallbackID: callbackID}, nil
}
