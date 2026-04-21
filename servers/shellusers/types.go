package shellusers

import "github.com/webdock-io/go-sdk/client"

type PublicKeyDTO struct {
	ID      int64  `json:"id" tfsdk:"id"`
	Name    string `json:"name" tfsdk:"name"`
	Key     string `json:"key" tfsdk:"key"`
	Created string `json:"created" tfsdk:"created"`
}

type ShellUser struct {
	ID         int64          `json:"id" tfsdk:"id"`
	Username   string         `json:"username" tfsdk:"username"`
	Group      string         `json:"group" tfsdk:"group"`
	Shell      string         `json:"shell" tfsdk:"shell"`
	PublicKeys []PublicKeyDTO `json:"publicKeys" tfsdk:"public_keys"`
	Created    string         `json:"created" tfsdk:"created"`
}

type ShellUsers struct {
	client *client.Client
}

type UpdatedShellUser struct {
	ShellUser  ShellUser `json:"shell_user" tfsdk:"shell_user"`
	CallbackID string    `json:"callbackId" tfsdk:"callback_id"`
}

func New(c *client.Client) ShellUsers {
	return ShellUsers{client: c}
}
