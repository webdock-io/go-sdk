package serverscripts

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/webdock-io/go-sdk/client"
)

type CreateScriptOptions struct {
	ServerSlug           string `json:"-"`
	ScriptId             int    `json:"scriptId"`
	Path                 string `json:"path"`
	MakeScriptExecutable bool   `json:"makeScriptExecutable,omitempty"`
	ExecuteImmediately   bool   `json:"executeImmediately,omitempty"`
}

type CreateScriptResponse struct {
	Script     Script `tfsdk:"script"`
	CallbackID string `tfsdk:"callback_id"`
}

func (s *ServerScripts) Create(opts CreateScriptOptions) (*CreateScriptResponse, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	var out Script
	c, err := s.client.Do("POST", fmt.Sprintf("v1/servers/%s/scripts", opts.ServerSlug), bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &CreateScriptResponse{Script: out, CallbackID: callbackID}, nil
}
