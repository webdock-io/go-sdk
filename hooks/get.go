package hooks

import "fmt"

type GetEventHookOptions struct {
	HookID int64
}

func (s *Hooks) GetByID(opts GetEventHookOptions) (*EventHookDTO, error) {
	var out EventHookDTO
	_, err := s.client.Do("GET", fmt.Sprintf("v1/hooks/%d", opts.HookID), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
