package serverscripts

import "github.com/webdock-io/go-sdk/client"

type Script struct {
	ID          int64  `json:"id" tfsdk:"id"`
	Name        string `json:"name" tfsdk:"name"`
	Description string `json:"description" tfsdk:"description"`
	Filename    string `json:"filename" tfsdk:"filename"`
	Content     string `json:"content" tfsdk:"content"`
}

type ServerScriptDTO struct {
	ID                int64  `json:"id" tfsdk:"id"`
	Name              string `json:"name" tfsdk:"name"`
	Path              string `json:"path" tfsdk:"path"`
	LastRun           string `json:"lastRun" tfsdk:"last_run"`
	LastRunCallbackId string `json:"lastRunCallbackId" tfsdk:"last_run_callback_id"`
	Created           string `json:"created" tfsdk:"created"`
}

type ServerScripts struct {
	client *client.Client
}

func New(c *client.Client) ServerScripts {
	return ServerScripts{client: c}
}
