// Package ftp implements Servers.FTP (src/servers/ftp in the Node SDK).
package ftp

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type User struct {
	ID       int64  `json:"id" tfsdk:"id"`
	Username string `json:"username" tfsdk:"username"`
	Group    string `json:"group" tfsdk:"group"`
	HomeDir  string `json:"homeDir" tfsdk:"home_dir"`
	Created  string `json:"created" tfsdk:"created"`
}

type DaemonStatus struct {
	Enabled    bool `json:"enabled" tfsdk:"enabled"`
	Manageable bool `json:"manageable" tfsdk:"manageable"`
}

type FTP struct {
	Users  Users
	Daemon Daemon
}

func New(c *client.Client) FTP {
	return FTP{Users: Users{client: c}, Daemon: Daemon{client: c}}
}

type ServerOptions struct {
	ServerSlug string
}

// Users implements servers.ftp.users.
type Users struct {
	client *client.Client
}

// List mirrors servers.ftp.users.list() in the Node SDK.
func (u *Users) List(ctx context.Context, opts ServerOptions) ([]User, error) {
	var out []User
	if _, err := u.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/ftp/users", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return out, nil
}

type CreateUserOptions struct {
	ServerSlug string `json:"-"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Group      string `json:"group,omitempty"`
	HomeDir    string `json:"homeDir,omitempty"`
}

// Create mirrors servers.ftp.users.create() in the Node SDK.
func (u *Users) Create(ctx context.Context, opts CreateUserOptions) (*client.AsyncActionResponse, error) {
	return u.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/ftp/users", opts.ServerSlug), opts)
}

type RemoveUserOptions struct {
	ServerSlug string
	FtpUserID  int64
}

// Remove mirrors servers.ftp.users.remove() in the Node SDK.
func (u *Users) Remove(ctx context.Context, opts RemoveUserOptions) (*client.AsyncActionResponse, error) {
	return u.client.DoAsync(ctx, "DELETE", fmt.Sprintf("/servers/%s/ftp/users/%d", opts.ServerSlug, opts.FtpUserID), nil)
}

// Daemon implements servers.ftp.daemon.
type Daemon struct {
	client *client.Client
}

// Get mirrors servers.ftp.daemon.get() in the Node SDK.
func (d *Daemon) Get(ctx context.Context, opts ServerOptions) (*DaemonStatus, error) {
	var out DaemonStatus
	if _, err := d.client.DoJSON(ctx, "GET", fmt.Sprintf("/servers/%s/ftp/daemon", opts.ServerSlug), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// Enable mirrors servers.ftp.daemon.enable() in the Node SDK.
func (d *Daemon) Enable(ctx context.Context, opts ServerOptions) (*client.AsyncActionResponse, error) {
	return d.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/ftp/daemon/enable", opts.ServerSlug), nil)
}

// Disable mirrors servers.ftp.daemon.disable() in the Node SDK.
func (d *Daemon) Disable(ctx context.Context, opts ServerOptions) (*client.AsyncActionResponse, error) {
	return d.client.DoAsync(ctx, "POST", fmt.Sprintf("/servers/%s/ftp/daemon/disable", opts.ServerSlug), nil)
}
