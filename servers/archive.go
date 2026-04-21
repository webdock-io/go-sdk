package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type ArchiveServerOptions struct {
	Slug string
}

func (s *Servers) Archive(ctx context.Context, opts ArchiveServerOptions) (string, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/suspend", opts.Slug), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
