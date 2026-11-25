// Package wordpress implements Servers.WordPress (src/servers/wordpress in
// the Node SDK).
package wordpress

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type UpdateType string

const (
	UpdateCore    UpdateType = "core"
	UpdatePlugins UpdateType = "plugins"
	UpdateAll     UpdateType = "all"
)

type Installation struct {
	Status string  `json:"status" tfsdk:"status"`
	Path   string  `json:"path" tfsdk:"path"`
	URL    *string `json:"url" tfsdk:"url"`
}

type WordPress struct {
	client       *client.Client
	Installation InstallationAPI
	Lockdown     Lockdown
	Updates      Updates
}

func New(c *client.Client) WordPress {
	return WordPress{
		client:       c,
		Installation: InstallationAPI{client: c},
		Lockdown:     Lockdown{client: c},
		Updates:      Updates{client: c},
	}
}

type ServerOptions struct {
	ServerSlug string
}

type SetupCLIOptions struct {
	ServerSlug string `json:"-"`
	// Path is the absolute path to an existing WordPress installation.
	Path string `json:"path,omitempty"`
	// URL is used when linking an existing installation through Path.
	URL client.Nullable[string] `json:"url,omitzero"`
}

// SetupCLI installs WP-CLI on the server. No body is sent unless Path or URL
// is set. Mirrors servers.wordpress.setupCli() in the Node SDK.
func (w *WordPress) SetupCLI(ctx context.Context, opts SetupCLIOptions) (*client.AsyncActionResponse, error) {
	var body any
	if opts.Path != "" || !opts.URL.IsZero() {
		body = opts
	}
	return w.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/wordpress/setup-cli", opts.ServerSlug), body)
}

// InstallationAPI implements servers.wordpress.installation.
type InstallationAPI struct {
	client *client.Client
}

type RegisterOptions struct {
	ServerSlug string                  `json:"-"`
	Path       string                  `json:"path"`
	URL        client.Nullable[string] `json:"url,omitzero"`
}

// Register links an existing WordPress installation. Mirrors
// servers.wordpress.installation.register() in the Node SDK.
func (i *InstallationAPI) Register(ctx context.Context, opts RegisterOptions) (*Installation, error) {
	var out Installation
	if _, err := i.client.DoJSON(ctx, "PATCH", fmt.Sprintf("/servers/%s/wordpress/installation", opts.ServerSlug), opts, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// InstallOptions installs a new WordPress site. Set DBUser and DBPass
// together, or leave both empty.
type InstallOptions struct {
	ServerSlug    string `json:"-"`
	Locale        string `json:"locale"`
	Path          string `json:"path"`
	URL           string `json:"url"`
	Title         string `json:"title"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	DBName        string `json:"dbName,omitempty"`
	DBHost        string `json:"dbHost,omitempty"`
	DBUser        string `json:"dbUser,omitempty"`
	DBPass        string `json:"dbPass,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
}

// Install mirrors servers.wordpress.installation.install() in the Node SDK.
func (i *InstallationAPI) Install(ctx context.Context, opts InstallOptions) (*client.AsyncActionResponse, error) {
	if (opts.DBUser == "") != (opts.DBPass == "") {
		return nil, fmt.Errorf("wordpress install: DBUser and DBPass must be set together")
	}
	return i.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/wordpress/install", opts.ServerSlug), opts)
}

// Lockdown implements servers.wordpress.lockdown.
type Lockdown struct {
	client *client.Client
}

type LockdownOptions struct {
	ServerSlug        string   `json:"-"`
	AllowedWritePaths []string `json:"allowedWritePaths,omitzero"`
	DisallowPHPPaths  []string `json:"disallowPhpPaths,omitzero"`
	WebUserGroup      string   `json:"webUserGroup,omitempty"`
}

// Enable locks down the WordPress installation. No body is sent unless an
// option is set. Mirrors servers.wordpress.lockdown.enable() in the Node SDK.
func (l *Lockdown) Enable(ctx context.Context, opts LockdownOptions) (*client.AsyncActionResponse, error) {
	var body any
	if opts.AllowedWritePaths != nil || opts.DisallowPHPPaths != nil || opts.WebUserGroup != "" {
		body = opts
	}
	return l.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/wordpress/lockdown", opts.ServerSlug), body)
}

// Disable mirrors servers.wordpress.lockdown.disable() in the Node SDK.
func (l *Lockdown) Disable(ctx context.Context, opts ServerOptions) (*client.AsyncActionResponse, error) {
	return l.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/wordpress/unlock", opts.ServerSlug), nil)
}

// Updates implements servers.wordpress.updates.
type Updates struct {
	client *client.Client
}

// Check mirrors servers.wordpress.updates.check() in the Node SDK.
func (u *Updates) Check(ctx context.Context, opts ServerOptions) (*client.AsyncActionResponse, error) {
	return u.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/wordpress/check-updates", opts.ServerSlug), nil)
}

type ApplyOptions struct {
	ServerSlug string     `json:"-"`
	UpdateType UpdateType `json:"updateType"`
}

// Apply mirrors servers.wordpress.updates.apply() in the Node SDK.
func (u *Updates) Apply(ctx context.Context, opts ApplyOptions) (*client.AsyncActionResponse, error) {
	return u.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/wordpress/update", opts.ServerSlug), opts)
}
