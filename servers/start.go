package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type StartServerOptions struct {
	Slug string
}

func (s *Servers) Start(ctx context.Context, opts StartServerOptions) (string, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/start", opts.Slug), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
