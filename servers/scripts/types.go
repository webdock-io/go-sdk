package serverscripts

import "github.com/webdock-io/go-sdk/client"

type Script struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Content     string `json:"content"`
}

type ServerScriptDTO struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	Path              string `json:"path"`
	LastRun           string `json:"lastRun"`
	LastRunCallbackId string `json:"lastRunCallbackId"`
	Created           string `json:"created"`
}

type ServerScripts struct {
	client *client.Client
}

func New(c *client.Client) ServerScripts {
	return ServerScripts{client: c}
}
