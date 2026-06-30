package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateServerFromImageOptions struct {
	Name           string `json:"name"`
	LocationId     string `json:"locationId"`
	ProfileSlug    string `json:"profileSlug,omitempty"`
	ImageSlug      string `json:"imageSlug,omitempty"`
	Virtualization string `json:"virtualization,omitempty"`
	UserScriptID   int64  `json:"userScriptId,omitempty"`
	UserScriptSlug string `json:"-"`
	Slug           string `json:"slug,omitempty"`
}

type CreateServerFromSnapshotOptions struct {
	Name           string `json:"name"`
	LocationId     string `json:"locationId"`
	ProfileSlug    string `json:"profileSlug,omitempty"`
	SnapshotId     int    `json:"snapshotId"`
	Virtualization string `json:"virtualization,omitempty"`
	UserScriptID   int64  `json:"userScriptId,omitempty"`
	UserScriptSlug string `json:"-"`
	Slug           string `json:"slug,omitempty"`
}

type CreateServerOptions struct {
	Name           string `json:"name"`
	LocationId     string `json:"locationId"`
	ProfileSlug    string `json:"profileSlug,omitempty"`
	ImageSlug      string `json:"imageSlug,omitempty"`
	SnapshotId     *int   `json:"snapshotId,omitempty"`
	Virtualization string `json:"virtualization,omitempty"`
	UserScriptID   int64  `json:"userScriptId,omitempty"`
	UserScriptSlug string `json:"-"`
	Slug           string `json:"slug,omitempty"`
}

type CreatedServer struct {
	Server     Server `json:"server" tfsdk:"server"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) Create(ctx context.Context, opts CreateServerOptions) (*CreatedServer, error) {
	data, err := json.Marshal(createServerPayload(opts))
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Server
	c, err := s.client.Do(ctx, "POST", "/servers", bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreatedServer{Server: out, CallbackID: callbackID}, nil
}

func createServerPayload(opts CreateServerOptions) map[string]any {
	payload := map[string]any{
		"name":       opts.Name,
		"locationId": opts.LocationId,
	}
	if opts.ProfileSlug != "" {
		payload["profileSlug"] = opts.ProfileSlug
	}
	if opts.ImageSlug != "" {
		payload["imageSlug"] = opts.ImageSlug
	}
	if opts.SnapshotId != nil {
		payload["snapshotId"] = *opts.SnapshotId
	}
	if opts.Virtualization != "" {
		payload["virtualization"] = opts.Virtualization
	}
	if opts.UserScriptSlug != "" {
		payload["userScriptId"] = opts.UserScriptSlug
	} else if opts.UserScriptID != 0 {
		payload["userScriptId"] = opts.UserScriptID
	}
	if opts.Slug != "" {
		payload["slug"] = opts.Slug
	}
	return payload
}

func (s *Servers) CreateFromImage(ctx context.Context, opts CreateServerFromImageOptions) (*CreatedServer, error) {
	return s.Create(ctx, CreateServerOptions{
		Name:           opts.Name,
		LocationId:     opts.LocationId,
		ProfileSlug:    opts.ProfileSlug,
		ImageSlug:      opts.ImageSlug,
		Virtualization: opts.Virtualization,
		UserScriptID:   opts.UserScriptID,
		UserScriptSlug: opts.UserScriptSlug,
		Slug:           opts.Slug,
	})
}

func (s *Servers) CreateFromSnapshot(ctx context.Context, opts CreateServerFromSnapshotOptions) (*CreatedServer, error) {
	return s.Create(ctx, CreateServerOptions{
		Name:           opts.Name,
		LocationId:     opts.LocationId,
		ProfileSlug:    opts.ProfileSlug,
		SnapshotId:     &opts.SnapshotId,
		Virtualization: opts.Virtualization,
		UserScriptID:   opts.UserScriptID,
		UserScriptSlug: opts.UserScriptSlug,
		Slug:           opts.Slug,
	})
}
