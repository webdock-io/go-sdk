package snapshots

import "context"

func (s *Snapshots) ListAll(ctx context.Context) ([]Snapshot, error) {
	var out []Snapshot
	_, err := s.client.Do(ctx, "GET", "/servers/snapshots", nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
