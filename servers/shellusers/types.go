package shellusers

import "github.com/webdock-io/go-sdk/client"

type PublicKeyDTO struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Created string `json:"created"`
}

type ShellUser struct {
	ID         int64          `json:"id"`
	Username   string         `json:"username"`
	Group      string         `json:"group"`
	Shell      string         `json:"shell"`
	PublicKeys []PublicKeyDTO `json:"publicKeys"`
	Created    string         `json:"created"`
}

type ShellUsers struct {
	client *client.Client
}

func New(c *client.Client) ShellUsers {
	return ShellUsers{client: c}
}
