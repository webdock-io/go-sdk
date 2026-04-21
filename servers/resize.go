package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ResizeServerOptions struct {
	Slug        string
	ProfileSlug string
}

func (s *Servers) Resize(ctx context.Context, opts ResizeServerOptions) (string, error) {
	data, err := json.Marshal(map[string]string{"profileSlug": opts.ProfileSlug})
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/resize", opts.Slug), bytes.NewBuffer(data), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
