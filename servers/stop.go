package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type StopServerOptions struct {
	Slug string
}

func (s *Servers) Stop(ctx context.Context, opts StopServerOptions) (string, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/stop", opts.Slug), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
