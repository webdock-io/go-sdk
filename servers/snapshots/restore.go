package snapshots

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type RestoreSnapshotOptions struct {
	ServerSlug string
	SnapshotID int64
}

type RestoreSnapshotResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Snapshots) Restore(ctx context.Context, opts RestoreSnapshotOptions) (*RestoreSnapshotResponse, error) {
	body, err := json.Marshal(map[string]int64{"snapshotId": opts.SnapshotID})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}

	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/actions/restore", opts.ServerSlug), bytes.NewBuffer(body), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &RestoreSnapshotResponse{CallbackID: callbackID}, nil
}
