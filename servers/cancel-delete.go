package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CancelDeleteOptions struct {
	ServerSlug string
}

func (s *Servers) CancelDelete(ctx context.Context, opts CancelDeleteOptions) (string, error) {
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/uncancel", opts.ServerSlug), nil, nil)
	if err != nil {
		return "", err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return callbackID, nil
}
