package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ReinstallServerOptions struct {
	Slug            string `json:"-"`
	ImageSlug       string `json:"imageSlug"`
	UserScriptID    int64  `json:"userScriptId,omitempty"`
	UserScriptSlug  string `json:"-"`
	DeleteSnapshots bool   `json:"deleteSnapshots,omitempty"`
}

type ReinstallServerResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Reinstall(ctx context.Context, opts ReinstallServerOptions) (*ReinstallServerResponse, error) {
	data, err := json.Marshal(reinstallServerPayload(opts))
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/reinstall", opts.Slug), bytes.NewBuffer(data), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &ReinstallServerResponse{CallbackID: callbackID}, nil
}

func reinstallServerPayload(opts ReinstallServerOptions) map[string]any {
	payload := map[string]any{
		"imageSlug": opts.ImageSlug,
	}
	if opts.DeleteSnapshots {
		payload["deleteSnapshots"] = opts.DeleteSnapshots
	}
	if opts.UserScriptSlug != "" {
		payload["userScriptId"] = opts.UserScriptSlug
	} else if opts.UserScriptID != 0 {
		payload["userScriptId"] = opts.UserScriptID
	}
	return payload
}
