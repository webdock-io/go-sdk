package hooks

import "context"

type DeleteByIDOptions struct {
	ID int64
}

func (s *Hooks) DeleteByID(ctx context.Context, opts DeleteByIDOptions) error {
	return s.Delete(ctx, DeleteEventHookOptions{HookID: opts.ID})
}
