package servers

import (
	"github.com/webdock-io/go-sdk/client"
	serverscripts "github.com/webdock-io/go-sdk/servers/scripts"
	"github.com/webdock-io/go-sdk/servers/shellusers"
	"github.com/webdock-io/go-sdk/servers/snapshots"
)

type ServerState string

const (
	Provisioning ServerState = "provisioning"
	Running      ServerState = "running"
	Stopped      ServerState = "stopped"
	Error        ServerState = "error"
	Rebooting    ServerState = "rebooting"
	Starting     ServerState = "starting"
	Stopping     ServerState = "stopping"
	Reinstalling ServerState = "reinstalling"
	Suspended    ServerState = "suspended"
)

type Virtualization string

const (
	Container Virtualization = "container"
	KVM       Virtualization = "kvm"
)

type ServerType string

const (
	Apache ServerType = "Apache"
	Nginx  ServerType = "Nginx"
	None   ServerType = "None"
)

type Server struct {
	Slug                   string         `json:"slug"`
	Name                   string         `json:"name"`
	Date                   string         `json:"date"`
	Location               string         `json:"location"`
	Image                  string         `json:"image"`
	Profile                string         `json:"profile"`
	IPv4                   string         `json:"ipv4"`
	IPv6                   string         `json:"ipv6"`
	Status                 ServerState    `json:"status"`
	Virtualization         Virtualization `json:"virtualization"`
	WebServer              ServerType     `json:"webServer"`
	Aliases                []string       `json:"aliases"`
	SnapshotRunTime        int            `json:"snapshotRunTime"`
	Description            string         `json:"description"`
	WordPressLockDown      bool           `json:"WordPressLockDown"`
	SSHPasswordAuthEnabled bool           `json:"SSHPasswordAuthEnabled"`
	Notes                  string         `json:"notes"`
	NextActionDate         string         `json:"nextActionDate"`
}

type Servers struct {
	client     *client.Client
	Scripts    serverscripts.ServerScripts
	ShellUsers shellusers.ShellUsers
	Snapshots  snapshots.Snapshots
}

func New(c *client.Client) Servers {
	return Servers{
		client:     c,
		Scripts:    serverscripts.New(c),
		ShellUsers: shellusers.New(c),
		Snapshots:  snapshots.New(c),
	}
}
