package sshkeys

import "github.com/webdock-io/go-sdk/client"

type SSHKey struct {
	ID      int64  `json:"id" tfsdk:"id"`
	Name    string `json:"name" tfsdk:"name"`
	Key     string `json:"key" tfsdk:"key"`
	Created string `json:"created" tfsdk:"created"`
}

type SSHSettings struct {
	ServerSlug              string `json:"-" tfsdk:"server_slug"`
	PasswordSSHAuthEnabled  *bool  `json:"passwordSshAuthEnabled,omitempty" tfsdk:"password_ssh_auth_enabled"`
	PasswordlessSudoEnabled *bool  `json:"passwordlessSudoEnabled,omitempty" tfsdk:"passwordless_sudo_enabled"`
	SSHPort                 *int   `json:"sshPort,omitempty" tfsdk:"ssh_port"`
}

type SSHKeys struct {
	client *client.Client
}

func New(c *client.Client) SSHKeys {
	return SSHKeys{client: c}
}
