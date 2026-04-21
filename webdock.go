package webdock

import (
	"github.com/webdock-io/go-sdk/account"
	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/events"
	"github.com/webdock-io/go-sdk/hooks"
	"github.com/webdock-io/go-sdk/images"
	"github.com/webdock-io/go-sdk/loactions"
	"github.com/webdock-io/go-sdk/locations"
	"github.com/webdock-io/go-sdk/operation"
	"github.com/webdock-io/go-sdk/platforms"
	"github.com/webdock-io/go-sdk/profiles"
	"github.com/webdock-io/go-sdk/servers"
	"github.com/webdock-io/go-sdk/servers/shellusers"
	"github.com/webdock-io/go-sdk/servers/snapshots"
	"github.com/webdock-io/go-sdk/sshkeys"
	platform "github.com/webdock-io/go-sdk/webdock"
	"github.com/webdock-io/go-sdk/webssh"
)

type Webdock struct {
	Account    account.Account
	Events     events.Events
	Hooks      hooks.Hooks
	Images     images.Images
	Location   locations.Locations
	Locations  loactions.Locations
	Operation  operation.Operation
	Platforms  platforms.Platforms
	Profiles   profiles.Profiles
	ShellUsers shellusers.ShellUsers
	SSHKeys    sshkeys.SSHKeys
	Snapshots  snapshots.Snapshots
	Servers    servers.Servers
	Webdock    platform.WebdockPlatform
	Webssh     webssh.Webssh
}

func New(token string) Webdock {
	c := client.New(token)
	return Webdock{
		Account:    account.New(c),
		Events:     events.New(c),
		Hooks:      hooks.New(c),
		Images:     images.New(c),
		Location:   locations.New(c),
		Locations:  loactions.New(c),
		Operation:  operation.New(c),
		Platforms:  platforms.New(c),
		Profiles:   profiles.New(c),
		ShellUsers: shellusers.New(c),
		SSHKeys:    sshkeys.New(c),
		Snapshots:  snapshots.New(c),
		Servers:    servers.New(c),
		Webdock:    platform.New(c),
		Webssh:     webssh.New(c),
	}
}
