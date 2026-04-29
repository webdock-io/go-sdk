package sshkeys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ToggleSSHSettingsResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *SSHKeys) ToggleSSHSettings(ctx context.Context, opts SSHSettings) (*ToggleSSHSettingsResponse, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/sshSettings", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, err
	}

	callbackID, _ := c.GetHeader(client.CallbackID)
	return &ToggleSSHSettingsResponse{CallbackID: callbackID}, nil
}
