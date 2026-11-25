package serverscripts

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateScriptOptions struct {
	ServerSlug           string `json:"-"`
	ScriptId             int    `json:"-"`
	ScriptSlug           string `json:"-"`
	Path                 string `json:"path"`
	MakeScriptExecutable bool   `json:"makeScriptExecutable"`
	ExecuteImmediately   bool   `json:"executeImmediately"`
}

type CreateScriptResponse struct {
	Script     Script `json:"script" tfsdk:"script"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// Create mirrors servers.scripts.create() in the Node SDK.
func (s *ServerScripts) Create(ctx context.Context, opts CreateScriptOptions) (*CreateScriptResponse, error) {
	data, err := json.Marshal(createScriptPayload(opts))
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Script
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("/servers/%s/scripts", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreateScriptResponse{Script: out, CallbackID: callbackID}, nil
}

func createScriptPayload(opts CreateScriptOptions) map[string]any {
	// The Node SDK always sends both flags, so false is sent explicitly too.
	payload := map[string]any{
		"path":                 opts.Path,
		"makeScriptExecutable": opts.MakeScriptExecutable,
		"executeImmediately":   opts.ExecuteImmediately,
	}
	if opts.ScriptSlug != "" {
		payload["scriptId"] = opts.ScriptSlug
	} else {
		payload["scriptId"] = opts.ScriptId
	}
	return payload
}
