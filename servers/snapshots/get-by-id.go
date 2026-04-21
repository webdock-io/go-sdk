package snapshots

import (
	"context"
	"fmt"
)

type GetSnapshotByIDOptions struct {
	ServerSlug string
	SnapshotId int64
}

func (s *Snapshots) GetByID(ctx context.Context, opts GetSnapshotByIDOptions) (*Snapshot, error) {
	var out Snapshot
	_, err := s.client.Do(ctx, "GET", fmt.Sprintf("v1/servers/%s/snapshots/%d", opts.ServerSlug, opts.SnapshotId), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
