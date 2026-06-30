package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type UpdateServerOptions struct {
	ServerSlug     string `json:"-"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Notes          string `json:"notes,omitempty"`
	NextActionDate string `json:"nextActionDate,omitempty"`
}

func (s *Servers) Update(ctx context.Context, opts UpdateServerOptions) (*Server, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Server
	_, err = s.client.Do(ctx, "PATCH", fmt.Sprintf("v1/servers/%s", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
