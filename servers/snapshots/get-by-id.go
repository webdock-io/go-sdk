package snapshots

import "fmt"

type GetSnapshotByIDOptions struct {
	ServerSlug string
	SnapshotId int64
}

func (s *Snapshots) GetByID(opts GetSnapshotByIDOptions) (*Snapshot, error) {
	var out Snapshot
	_, err := s.client.Do("GET", fmt.Sprintf("v1/servers/%s/snapshots/%d", opts.ServerSlug, opts.SnapshotId), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
