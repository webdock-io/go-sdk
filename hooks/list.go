package hooks

type ListEventHooksOptions struct{}

type ListEventHooksResponse []EventHookDTO

func (s *Hooks) List(opts ListEventHooksOptions) (*ListEventHooksResponse, error) {
	var out ListEventHooksResponse
	_, err := s.client.Do("GET", "v1/hooks", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
