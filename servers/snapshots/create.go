package snapshots

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type TakeSnapshotOptions struct {
	ServerSlug string
	Name       string
}

type TakeSnapshotResponse struct {
	Snapshot   Snapshot `json:"snapshot" tfsdk:"snapshot"`
	CallbackID string   `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Snapshots) Take(ctx context.Context, opts TakeSnapshotOptions) (*TakeSnapshotResponse, error) {
	data, err := json.Marshal(map[string]string{"name": opts.Name})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Snapshot
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/snapshot", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &TakeSnapshotResponse{Snapshot: out, CallbackID: callbackID}, nil
}

func (s *Snapshots) Create(ctx context.Context, opts TakeSnapshotOptions) (*TakeSnapshotResponse, error) {
	return s.Take(ctx, opts)
}
