package hooks

import (
	"context"
	"fmt"
)

type GetEventHookOptions struct {
	HookID int64
}

func (s *Hooks) GetByID(ctx context.Context, opts GetEventHookOptions) (*EventHookDTO, error) {
	var out EventHookDTO
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("v1/hooks/%d", opts.HookID), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
