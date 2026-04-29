package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type RestoreFromSnapshotOptions struct {
	ServerSlug string
	SnapshotId string
}

type RestoreFromSnapshotResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) RestoreFromSnapshot(ctx context.Context, opts RestoreFromSnapshotOptions) (*RestoreFromSnapshotResponse, error) {
	data, err := json.Marshal(map[string]string{"snapshotId": opts.SnapshotId})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/actions/restore", opts.ServerSlug), bytes.NewBuffer(data), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &RestoreFromSnapshotResponse{CallbackID: callbackID}, nil
}
