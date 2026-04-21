package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ReinstallServerOptions struct {
	Slug      string
	ImageSlug string
}

func (s *Servers) Reinstall(ctx context.Context, opts ReinstallServerOptions) (string, error) {
	data, err := json.Marshal(map[string]string{"imageSlug": opts.ImageSlug})
	if err != nil {
		return "", fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/reinstall", opts.Slug), bytes.NewBuffer(data), nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
