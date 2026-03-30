package sdk

import (
	"github.com/webdock-io/go-sdk/account"
	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/events"
	"github.com/webdock-io/go-sdk/hooks"
	"github.com/webdock-io/go-sdk/images"
	"github.com/webdock-io/go-sdk/loactions"
	"github.com/webdock-io/go-sdk/profiles"
	"github.com/webdock-io/go-sdk/servers"
	"github.com/webdock-io/go-sdk/webdock"
	"github.com/webdock-io/go-sdk/webssh"
)

type Webdock struct {
	client    *client.Client
	Account   account.Account
	Events    events.Events
	Hooks     hooks.Hooks
	Images    images.Images
	Locations loactions.Locations
	Profiles  profiles.Profiles
	Servers   servers.Servers
	Webdock   webdock.Webdock
	Webssh    webssh.Webssh
}

func New(token string) Webdock {
	c := client.New(token)
	return Webdock{
		client:    c,
		Account:   account.New(c),
		Events:    events.New(c),
		Hooks:     hooks.New(c),
		Images:    images.New(c),
		Locations: loactions.New(c),
		Profiles:  profiles.New(c),
		Servers:   servers.New(c),
		Webdock:   webdock.New(c),
		Webssh:    webssh.New(c),
	}
}
