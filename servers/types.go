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
	Slug                   string         `json:"slug" tfsdk:"slug"`
	Name                   string         `json:"name" tfsdk:"name"`
	Date                   string         `json:"date" tfsdk:"date"`
	Location               string         `json:"location" tfsdk:"location"`
	Image                  string         `json:"image" tfsdk:"image"`
	Profile                string         `json:"profile" tfsdk:"profile"`
	IPv4                   string         `json:"ipv4" tfsdk:"ipv4"`
	IPv6                   string         `json:"ipv6" tfsdk:"ipv6"`
	Status                 ServerState    `json:"status" tfsdk:"status"`
	Virtualization         Virtualization `json:"virtualization" tfsdk:"virtualization"`
	WebServer              ServerType     `json:"webServer" tfsdk:"web_server"`
	Aliases                []string       `json:"aliases" tfsdk:"aliases"`
	SnapshotRunTime        int            `json:"snapshotRunTime" tfsdk:"snapshot_run_time"`
	Description            string         `json:"description" tfsdk:"description"`
	WordPressLockDown      bool           `json:"WordPressLockDown" tfsdk:"wordpress_lock_down"`
	SSHPasswordAuthEnabled bool           `json:"SSHPasswordAuthEnabled" tfsdk:"ssh_password_auth_enabled"`
	Notes                  string         `json:"notes" tfsdk:"notes"`
	NextActionDate         string         `json:"nextActionDate" tfsdk:"next_action_date"`
}

type Servers struct {
	client     *client.Client
	Identity   ServerIdentity
	Settings   ServerSettings
	Scripts    serverscripts.ServerScripts
	ShellUsers shellusers.ShellUsers
	Snapshots  snapshots.Snapshots
}

func New(c *client.Client) Servers {
	return Servers{
		client:     c,
		Identity:   NewServerIdentity(c),
		Settings:   NewServerSettings(c),
		Scripts:    serverscripts.New(c),
		ShellUsers: shellusers.New(c),
		Snapshots:  snapshots.New(c),
	}
}
