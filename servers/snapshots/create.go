package snapshots

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type TakeSnapshotOptions struct {
	ServerSlug string
	Name       string
}

type TakeSnapshotResponse struct {
	Snapshot   Snapshot `tfsdk:"snapshot"`
	CallbackID string   `tfsdk:"callback_id"`
}

func (s *Snapshots) Take(opts TakeSnapshotOptions) (*TakeSnapshotResponse, error) {
	data, err := json.Marshal(map[string]string{"name": opts.Name})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Snapshot
	c, err := s.client.Do("POST", fmt.Sprintf("v1/servers/%s/actions/snapshot", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &TakeSnapshotResponse{Snapshot: out, CallbackID: callbackID}, nil
}
