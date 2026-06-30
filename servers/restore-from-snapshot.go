package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type RestoreFromSnapshotOptions struct {
	ServerSlug string
	SnapshotId string
	SnapshotID int64
}

type RestoreFromSnapshotResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) RestoreFromSnapshot(ctx context.Context, opts RestoreFromSnapshotOptions) (*RestoreFromSnapshotResponse, error) {
	data, err := json.Marshal(restoreFromSnapshotPayload(opts))
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

func restoreFromSnapshotPayload(opts RestoreFromSnapshotOptions) map[string]any {
	if opts.SnapshotID != 0 {
		return map[string]any{"snapshotId": opts.SnapshotID}
	}
	if opts.SnapshotId != "" {
		if id, err := strconv.ParseInt(opts.SnapshotId, 10, 64); err == nil {
			return map[string]any{"snapshotId": id}
		}
		return map[string]any{"snapshotId": opts.SnapshotId}
	}
	return map[string]any{"snapshotId": int64(0)}
}
