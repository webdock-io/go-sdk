package snapshots

import (
	"context"
	"fmt"
)

type ListSnapshotsOptions struct {
	ServerSlug string
}

func (s *Snapshots) List(ctx context.Context, opts ListSnapshotsOptions) ([]Snapshot, error) {
	var out []Snapshot
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("v1/servers/%s/snapshots", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
