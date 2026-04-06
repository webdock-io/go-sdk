package servers

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateServerFromImageOptions struct {
	Name        string `json:"name"`
	LocationId  string `json:"locationId"`
	ProfileSlug string `json:"profileSlug"`
	ImageSlug   string `json:"imageSlug"`
}

type CreateServerFromSnapshotOptions struct {
	Name        string `json:"name"`
	LocationId  string `json:"locationId"`
	ProfileSlug string `json:"profileSlug"`
	SnapshotId  int    `json:"snapshotId"`
}

type CreatedServer struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `tfsdk:"callback_id"`
}

func (s *Servers) CreateFromImage(opts CreateServerFromImageOptions) (*CreatedServer, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Server
	c, err := s.client.Do("POST", "v1/servers", bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreatedServer{Server: out, CallbackID: callbackID}, nil
}

func (s *Servers) CreateFromSnapshot(opts CreateServerFromSnapshotOptions) (*CreatedServer, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Server
	c, err := s.client.Do("POST", "v1/servers", bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreatedServer{Server: out, CallbackID: callbackID}, nil
}
