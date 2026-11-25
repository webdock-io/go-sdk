package mail

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type Mail struct {
	client     *client.Client
	Cloudflare Cloudflare
	Postmark   Postmark
	Improvmx   Improvmx
}

func New(c *client.Client) Mail {
	return Mail{
		client:     c,
		Cloudflare: Cloudflare{client: c, DefaultAlias: CloudflareDefaultAlias{client: c}},
		Postmark:   Postmark{client: c},
		Improvmx:   Improvmx{client: c, Aliases: ImprovmxAliases{client: c}},
	}
}

type ServerOptions struct {
	ServerSlug string
}

// Get mirrors servers.mail.get() in the Node SDK.
func (m *Mail) Get(ctx context.Context, opts ServerOptions) (*Management, error) {
	return getManagement(ctx, m.client, "GET", fmt.Sprintf("/servers/%s/mail", opts.ServerSlug), nil)
}

func getManagement(ctx context.Context, c *client.Client, method, path string, body any) (*Management, error) {
	var out Management
	if _, err := c.DoJSON(ctx, method, path, body, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Cloudflare implements servers.mail.cloudflare.
type Cloudflare struct {
	client       *client.Client
	DefaultAlias CloudflareDefaultAlias
}

// Options mirrors servers.mail.cloudflare.options() in the Node SDK.
func (c *Cloudflare) Options(ctx context.Context, opts ServerOptions) (*CloudflareOptions, error) {
	var out CloudflareOptions
	if _, err := c.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/mail/cloudflare/options", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

type ZonesOptions struct {
	ServerSlug string
	AccountID  int64
}

// Zones mirrors servers.mail.cloudflare.zones() in the Node SDK.
func (c *Cloudflare) Zones(ctx context.Context, opts ZonesOptions) ([]CloudflareZone, error) {
	var out []CloudflareZone
	if _, err := c.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/mail/cloudflare/accounts/%d/zones", opts.ServerSlug, opts.AccountID), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// CloudflareDefaultAlias implements servers.mail.cloudflare.defaultAlias.
type CloudflareDefaultAlias struct {
	client *client.Client
}

type UpdateDefaultAliasOptions struct {
	ServerSlug string `json:"-"`
	AccountID  int64  `json:"accountId"`
	ZoneID     string `json:"zoneId"`
	ZoneName   string `json:"zoneName"`
	Pattern    string `json:"pattern,omitempty"`
}

// Update mirrors servers.mail.cloudflare.defaultAlias.update() in the Node SDK.
func (d *CloudflareDefaultAlias) Update(ctx context.Context, opts UpdateDefaultAliasOptions) (*CloudflareOptions, error) {
	var out CloudflareOptions
	if _, err := d.client.DoJSON(ctx, "PATCH", fmt.Sprintf("/servers/%s/mail/cloudflare/default-alias", opts.ServerSlug), opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Clear mirrors servers.mail.cloudflare.defaultAlias.clear() in the Node SDK.
func (d *CloudflareDefaultAlias) Clear(ctx context.Context, opts ServerOptions) (*CloudflareOptions, error) {
	var out CloudflareOptions
	if _, err := d.client.DoJSON(ctx, "DELETE", fmt.Sprintf("/servers/%s/mail/cloudflare/default-alias", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Postmark implements servers.mail.postmark.
type Postmark struct {
	client *client.Client
}

// Enable mirrors servers.mail.postmark.enable() in the Node SDK.
func (p *Postmark) Enable(ctx context.Context, opts EnableOptions) (*Management, error) {
	return getManagement(ctx, p.client, "POST", fmt.Sprintf("/servers/%s/mail/postmark/enable", opts.ServerSlug), opts)
}

// Disable mirrors servers.mail.postmark.disable() in the Node SDK.
func (p *Postmark) Disable(ctx context.Context, opts ServerOptions) (*Management, error) {
	return getManagement(ctx, p.client, "POST", fmt.Sprintf("/servers/%s/mail/postmark/disable", opts.ServerSlug), nil)
}

// Improvmx implements servers.mail.improvmx.
type Improvmx struct {
	client  *client.Client
	Aliases ImprovmxAliases
}

// Enable mirrors servers.mail.improvmx.enable() in the Node SDK.
func (i *Improvmx) Enable(ctx context.Context, opts EnableOptions) (*Management, error) {
	return getManagement(ctx, i.client, "POST", fmt.Sprintf("/servers/%s/mail/improvmx/enable", opts.ServerSlug), opts)
}

// Disable mirrors servers.mail.improvmx.disable() in the Node SDK.
func (i *Improvmx) Disable(ctx context.Context, opts ServerOptions) (*Management, error) {
	return getManagement(ctx, i.client, "POST", fmt.Sprintf("/servers/%s/mail/improvmx/disable", opts.ServerSlug), nil)
}

// ImprovmxAliases implements servers.mail.improvmx.aliases.
type ImprovmxAliases struct {
	client *client.Client
}

// List mirrors servers.mail.improvmx.aliases.list() in the Node SDK.
func (a *ImprovmxAliases) List(ctx context.Context, opts ServerOptions) ([]ImprovmxAlias, error) {
	var out []ImprovmxAlias
	if _, err := a.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/mail/improvmx/aliases", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type CreateAliasOptions struct {
	ServerSlug  string `json:"-"`
	Alias       string `json:"alias"`
	Destination string `json:"destination"`
}

// Create mirrors servers.mail.improvmx.aliases.create() in the Node SDK.
func (a *ImprovmxAliases) Create(ctx context.Context, opts CreateAliasOptions) ([]ImprovmxAlias, error) {
	var out []ImprovmxAlias
	if _, err := a.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/mail/improvmx/aliases", opts.ServerSlug), opts, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type RemoveAliasOptions struct {
	ServerSlug string `json:"-"`
	Alias      string `json:"alias"`
}

// Remove mirrors servers.mail.improvmx.aliases.remove() in the Node SDK.
func (a *ImprovmxAliases) Remove(ctx context.Context, opts RemoveAliasOptions) ([]ImprovmxAlias, error) {
	var out []ImprovmxAlias
	if _, err := a.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/mail/improvmx/aliases/remove", opts.ServerSlug), opts, &out); err != nil {
		return nil, err
	}
	return out, nil
}
