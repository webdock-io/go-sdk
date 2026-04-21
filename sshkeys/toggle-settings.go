package sshkeys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

func (s *SSHKeys) ToggleSSHSettings(ctx context.Context, opts SSHSettings) (string, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/sshSettings", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return "", err
	}

	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
