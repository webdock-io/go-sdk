// Package webdock is the Go SDK for the Webdock API. Its API surface mirrors
// the Node SDK (@webdock/sdk) one to one: client.Servers.Mail.Improvmx.Aliases.List
// here is client.servers.mail.improvmx.aliases.list() there.
package webdock

import (
	"github.com/webdock-io/go-sdk/account"
	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/events"
	"github.com/webdock-io/go-sdk/hooks"
	"github.com/webdock-io/go-sdk/images"
	"github.com/webdock-io/go-sdk/loactions"
	"github.com/webdock-io/go-sdk/loadbalancers"
	"github.com/webdock-io/go-sdk/locations"
	"github.com/webdock-io/go-sdk/operation"
	"github.com/webdock-io/go-sdk/platforms"
	"github.com/webdock-io/go-sdk/profiles"
	"github.com/webdock-io/go-sdk/servers"
	"github.com/webdock-io/go-sdk/servers/shellusers"
	"github.com/webdock-io/go-sdk/servers/snapshots"
	"github.com/webdock-io/go-sdk/sshkeys"
	"github.com/webdock-io/go-sdk/wafs"
	platform "github.com/webdock-io/go-sdk/webdock"
	"github.com/webdock-io/go-sdk/webssh"
)

type Webdock struct {
	Account       account.Account
	Images        images.Images
	Profiles      profiles.Profiles
	Events        events.Events
	Hooks         hooks.Hooks
	Location      locations.Locations
	Servers       servers.Servers
	LoadBalancers loadbalancers.LoadBalancers
	Wafs          wafs.Wafs
	ShellUsers    shellusers.ShellUsers
	SSHKeys       sshkeys.SSHKeys
	Snapshots     snapshots.Snapshots
	Operation     operation.Operation
	Webdock       platform.WebdockPlatform
	Platforms     platforms.Platforms

	// Deprecated: use Location. This field comes from a misspelled package
	// ("loactions") and has no Node SDK counterpart.
	Locations loactions.Locations
	// Deprecated: use ShellUsers.WebSSHToken, which is where the Node SDK
	// exposes it.
	Webssh webssh.Webssh
}

// New returns a client authenticated with an API token.
func New(token string) Webdock {
	return NewWithClient(client.New(token))
}

// NewWithClient builds the SDK on top of an existing client, for example one
// created with client.NewWithBaseURL.
func NewWithClient(c *client.Client) Webdock {
	return Webdock{
		Account:       account.New(c),
		Images:        images.New(c),
		Profiles:      profiles.New(c),
		Events:        events.New(c),
		Hooks:         hooks.New(c),
		Location:      locations.New(c),
		Servers:       servers.New(c),
		LoadBalancers: loadbalancers.New(c),
		Wafs:          wafs.New(c),
		ShellUsers:    shellusers.New(c),
		SSHKeys:       sshkeys.New(c),
		Snapshots:     snapshots.New(c),
		Operation:     operation.New(c),
		Webdock:       platform.New(c),
		Platforms:     platforms.New(c),
		Locations:     loactions.New(c),
		Webssh:        webssh.New(c),
	}
}
