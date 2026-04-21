package hooks

import "context"

type ListEventHooksResponse []EventHookDTO

func (s *Hooks) List(ctx context.Context) (*ListEventHooksResponse, error) {
	var out ListEventHooksResponse
	_, err := s.client.Do(ctx, "GET", "/hooks", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
