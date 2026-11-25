package servers

import (
	"context"
	"fmt"
)

type ResizeServerOptions struct {
	Slug        string
	ProfileSlug string
}

type ResizeServerResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Resize mirrors servers.resize() in the Node SDK.
func (s *Servers) Resize(ctx context.Context, opts ResizeServerOptions) (*ResizeServerResponse, error) {
	var out Server
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/actions/resize", opts.Slug), map[string]string{"profileSlug": opts.ProfileSlug}, &out)
	if err != nil {
		return nil, err
	}
	return &ResizeServerResponse{Server: out, CallbackID: callbackID}, nil
}
