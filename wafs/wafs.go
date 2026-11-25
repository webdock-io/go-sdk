// Package wafs implements Webdock.Wafs (src/wafs in the Node SDK).
package wafs

import (
	"context"
	"net/url"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/edge"
)

type ProtectionLevel string

const (
	ProtectionMonitoring ProtectionLevel = "monitoring"
	ProtectionFull       ProtectionLevel = "full"
	ProtectionAdvanced   ProtectionLevel = "advanced"
)

type Waf struct {
	ID              string          `json:"id" tfsdk:"id"`
	Name            string          `json:"name" tfsdk:"name"`
	Domain          string          `json:"domain" tfsdk:"domain"`
	ProtectionLevel ProtectionLevel `json:"protectionLevel" tfsdk:"protection_level"`
	Upstreams       []edge.Upstream `json:"upstreams" tfsdk:"upstreams"`
	Settings        edge.Settings   `json:"settings" tfsdk:"settings"`
}

// Options are the WAF-specific toggles on top of edge.Options. Nil or empty
// fields are not sent.
type Options struct {
	edge.Options
	ProtectionLevel        ProtectionLevel `json:"protectionLevel,omitempty"`
	DDoSProtection         *bool           `json:"ddosProtection,omitempty"`
	DDoSProtectionAdvanced *bool           `json:"ddosProtectionAdvanced,omitempty"`
	BotguardProtection     *bool           `json:"botguardProtection,omitempty"`
}

type Wafs struct {
	client *client.Client
}

func New(c *client.Client) Wafs {
	return Wafs{client: c}
}

func path(id string) string {
	return "/wafs/" + url.PathEscape(id)
}

// List mirrors wafs.list() in the Node SDK.
func (w *Wafs) List(ctx context.Context) ([]Waf, error) {
	var out []Waf
	if _, err := w.client.DoJSON(ctx, "GET", "/wafs", nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type CreateOptions struct {
	Name      string               `json:"name"`
	Domain    string               `json:"domain"`
	Upstreams []edge.UpstreamInput `json:"upstreams"`
	Options
}

// Create mirrors wafs.create() in the Node SDK.
func (w *Wafs) Create(ctx context.Context, opts CreateOptions) (*Waf, error) {
	var out Waf
	if _, err := w.client.DoJSON(ctx, "POST", "/wafs", opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type GetByIDOptions struct {
	WafID string
}

// GetByID mirrors wafs.getById() in the Node SDK.
func (w *Wafs) GetByID(ctx context.Context, opts GetByIDOptions) (*Waf, error) {
	var out Waf
	if _, err := w.client.DoJSON(ctx, "GET", path(opts.WafID), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateOptions updates a WAF. Empty fields are not sent.
type UpdateOptions struct {
	WafID     string               `json:"-"`
	Name      string               `json:"name,omitempty"`
	Domain    string               `json:"domain,omitempty"`
	Upstreams []edge.UpstreamInput `json:"upstreams,omitzero"`
	Options
}

// Update mirrors wafs.update() in the Node SDK.
func (w *Wafs) Update(ctx context.Context, opts UpdateOptions) (*Waf, error) {
	var out Waf
	if _, err := w.client.DoJSON(ctx, "PATCH", path(opts.WafID), opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type RemoveOptions struct {
	WafID string
}

// Remove deletes a WAF (the API answers 204). Mirrors wafs.remove() in the
// Node SDK.
func (w *Wafs) Remove(ctx context.Context, opts RemoveOptions) error {
	_, err := w.client.DoJSON(ctx, "DELETE", path(opts.WafID), nil, nil)
	return err
}
