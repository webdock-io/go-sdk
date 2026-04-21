package profiles

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type CreateProfileOptions struct {
	Platform         Platform `json:"platform"`
	CPUThreads       int      `json:"cpu_threads"`
	RAM              int      `json:"ram"`
	DiskSpace        int      `json:"disk_space"`
	NetworkBandwidth int      `json:"network_bandwidth"`
}

func (p *Profiles) Create(ctx context.Context, opts CreateProfileOptions) (*Profile, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	var out Profile
	_, err = p.client.Do(ctx, "POST", "/profiles", bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
