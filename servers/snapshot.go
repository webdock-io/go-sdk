package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/servers/snapshots"
)

type SnapshotServerOptions struct {
	ServerSlug string
	Name       string
}

type SnapshotServerResponse struct {
	Snapshot   snapshots.Snapshot `json:"snapshot" tfsdk:"snapshot"`
	CallbackID string             `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Snapshot(ctx context.Context, opts SnapshotServerOptions) (*SnapshotServerResponse, error) {
	body, err := json.Marshal(map[string]string{"name": opts.Name})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	var out snapshots.Snapshot
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/snapshot", opts.ServerSlug), bytes.NewBuffer(body), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &SnapshotServerResponse{Snapshot: out, CallbackID: callbackID}, nil
}
