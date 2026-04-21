package hooks

import (
	"context"
	"fmt"
)

type DeleteEventHookOptions struct {
	HookID int64
}

func (s *Hooks) Delete(ctx context.Context, opts DeleteEventHookOptions) error {
	_, err := s.client.Do(ctx, "DELETE", fmt.Sprintf("v1/hooks/%d", opts.HookID), nil, nil)
	return err
}
