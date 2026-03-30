package hooks

import "fmt"

type DeleteEventHookOptions struct {
	HookID int64
}

func (s *Hooks) Delete(opts DeleteEventHookOptions) error {
	_, err := s.client.Do("DELETE", fmt.Sprintf("v1/hooks/%d", opts.HookID), nil, nil)
	return err
}
