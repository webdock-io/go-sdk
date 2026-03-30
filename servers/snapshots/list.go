package snapshots

import "fmt"

type ListSnapshotsOptions struct {
	ServerSlug string
}

func (s *Snapshots) List(opts ListSnapshotsOptions) ([]Snapshot, error) {
	var out []Snapshot
	_, err := s.client.Do("GET", fmt.Sprintf("v1/servers/%s/snapshots", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
