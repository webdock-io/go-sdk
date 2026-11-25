package servers

import (
	"context"
	"fmt"
)

type RebootServerOptions struct {
	Slug string
}

type RebootServerResponse struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Reboot mirrors servers.reboot() in the Node SDK.
func (s *Servers) Reboot(ctx context.Context, opts RebootServerOptions) (*RebootServerResponse, error) {
	var out Server
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/actions/reboot", opts.Slug), nil, &out)
	if err != nil {
		return nil, err
	}
	return &RebootServerResponse{Server: out, CallbackID: callbackID}, nil
}
