package servers

import (
	"encoding/json"

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

type ServerMetadata struct {
	DefaultAliasDisabled   bool   `json:"default_alias_disabled" tfsdk:"default_alias_disabled"`
	HasSetDomains          bool   `json:"has_set_domains" tfsdk:"has_set_domains"`
	CertbotLastRunTime     string `json:"certbot_lastrun_time" tfsdk:"certbot_lastrun_time"`
	CertbotLastRunResult   string `json:"certbot_lastrun_result" tfsdk:"certbot_lastrun_result"`
	Icon                   string `json:"icon" tfsdk:"icon"`
	Description            string `json:"description" tfsdk:"description"`
	Notes                  string `json:"notes" tfsdk:"notes"`
	InvoiceDate            string `json:"invoice_date" tfsdk:"invoice_date"`
	AutoStopOnBandwidthCap bool   `json:"auto_stop_on_bandwidth_cap" tfsdk:"auto_stop_on_bandwidth_cap"`
}

type ServerServices struct {
	IsManagedServer bool             `json:"is_managed_server" tfsdk:"is_managed_server"`
	ServiceList     []map[string]any `json:"service_list" tfsdk:"service_list"`
}

type ServerImage struct {
	Slug       string      `json:"slug" tfsdk:"slug"`
	Name       string      `json:"name" tfsdk:"name"`
	WebServer  *ServerType `json:"webServer" tfsdk:"web_server"`
	PHPVersion *string     `json:"phpVersion" tfsdk:"php_version"`
}

type ServerCPU struct {
	Cores   int `json:"cores" tfsdk:"cores"`
	Threads int `json:"threads" tfsdk:"threads"`
}

type ServerPrice struct {
	Amount   float64 `json:"amount" tfsdk:"amount"`
	Currency string  `json:"currency" tfsdk:"currency"`
}

type ServerProfile struct {
	Slug             string      `json:"slug" tfsdk:"slug"`
	Name             string      `json:"name" tfsdk:"name"`
	RAM              int         `json:"ram" tfsdk:"ram"`
	Disk             int         `json:"disk" tfsdk:"disk"`
	CPU              ServerCPU   `json:"cpu" tfsdk:"cpu"`
	Price            ServerPrice `json:"price" tfsdk:"price"`
	NetworkBandwidth int         `json:"network_bandwidth" tfsdk:"network_bandwidth"`
	Platform         string      `json:"platform" tfsdk:"platform"`
}

type Server struct {
	Slug                    string         `json:"slug" tfsdk:"slug"`
	Name                    string         `json:"name" tfsdk:"name"`
	Date                    string         `json:"date" tfsdk:"date"`
	Location                string         `json:"location" tfsdk:"location"`
	Image                   string         `json:"image" tfsdk:"image"`
	Profile                 string         `json:"profile" tfsdk:"profile"`
	IPv4                    string         `json:"ipv4" tfsdk:"ipv4"`
	IPv6                    string         `json:"ipv6" tfsdk:"ipv6"`
	Status                  ServerState    `json:"status" tfsdk:"status"`
	PendingDeletion         bool           `json:"pendingDeletion" tfsdk:"pending_deletion"`
	Virtualization          Virtualization `json:"virtualization" tfsdk:"virtualization"`
	WebServer               ServerType     `json:"webServer" tfsdk:"web_server"`
	Aliases                 []string       `json:"aliases" tfsdk:"aliases"`
	SnapshotRunTime         int            `json:"snapshotRunTime" tfsdk:"snapshot_run_time"`
	Description             string         `json:"description" tfsdk:"description"`
	WordPressLockDown       bool           `json:"WordPressLockDown" tfsdk:"wordpress_lock_down"`
	SSHPasswordAuthEnabled  bool           `json:"SSHPasswordAuthEnabled" tfsdk:"ssh_password_auth_enabled"`
	PasswordlessSudoEnabled bool           `json:"passwordlessSudoEnabled" tfsdk:"passwordless_sudo_enabled"`
	Notes                   string         `json:"notes" tfsdk:"notes"`
	NextActionDate          string         `json:"nextActionDate" tfsdk:"next_action_date"`
	Metadata                ServerMetadata `json:"metadata" tfsdk:"metadata"`
	Services                ServerServices `json:"services" tfsdk:"services"`
	SecondaryIPs            []string       `json:"secondaryIps" tfsdk:"secondary_ips"`
	LastChecked             string         `json:"lastChecked" tfsdk:"last_checked"`
	ImageData               ServerImage    `json:"imageData" tfsdk:"image_data"`
	ProfileData             ServerProfile  `json:"profileData" tfsdk:"profile_data"`
}

func (s *Server) UnmarshalJSON(data []byte) error {
	type serverAlias Server
	aux := struct {
		*serverAlias
		SecondaryIPsSnake []string `json:"secondary_ips"`
		LastCheckedLegacy string   `json:"lastchecked"`
	}{
		serverAlias: (*serverAlias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if len(s.SecondaryIPs) == 0 && aux.SecondaryIPsSnake != nil {
		s.SecondaryIPs = aux.SecondaryIPsSnake
	}
	if s.LastChecked == "" && aux.LastCheckedLegacy != "" {
		s.LastChecked = aux.LastCheckedLegacy
	}
	return nil
}

type Servers struct {
	client     *client.Client
	IPBlocks   ServerIPBlocks
	Identity   ServerIdentity
	Settings   ServerSettings
	Webserver  ServerWebserver
	Scripts    serverscripts.ServerScripts
	ShellUsers shellusers.ShellUsers
	Snapshots  snapshots.Snapshots
}

func New(c *client.Client) Servers {
	return Servers{
		client:     c,
		IPBlocks:   NewServerIPBlocks(c),
		Identity:   NewServerIdentity(c),
		Settings:   NewServerSettings(c),
		Webserver:  NewServerWebserver(c),
		Scripts:    serverscripts.New(c),
		ShellUsers: shellusers.New(c),
		Snapshots:  snapshots.New(c),
	}
}
