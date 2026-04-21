package servers

import (
	"context"
	"fmt"
)

type DeleteServerOptions struct {
	Slug string
}

func (s *Servers) Delete(ctx context.Context, opts DeleteServerOptions) error {
	_, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("v1/servers/%s", opts.Slug), nil, nil)
	return err
}
