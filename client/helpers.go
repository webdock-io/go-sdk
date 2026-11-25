package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// AsyncActionResponse is returned by endpoints that queue an action on a
// server. Body holds whatever JSON the API returned (often empty) and
// CallbackID is the X-Callback-ID header, which can be passed to
// Operation.WaitForEventToEnd or Events.List.
type AsyncActionResponse struct {
	Body       any    `json:"body" tfsdk:"body"`
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// DoJSON encodes body as JSON (when body is not nil), performs the request
// and decodes the response into out (when out is not nil). It returns the
// response's X-Callback-ID header, which is empty for synchronous endpoints.
func (c *Client) DoJSON(ctx context.Context, method, path string, body any, out any) (string, error) {
	var payload io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return "", fmt.Errorf("marshaling request: %w", err)
		}
		payload = bytes.NewReader(data)
	}

	res, err := c.Do(ctx, method, path, payload, out)
	if err != nil {
		return "", err
	}
	callbackID, _ := res.GetHeader(CallbackID)
	return callbackID, nil
}

// DoAsync performs a request against an endpoint that queues a server action
// and returns the raw body together with the X-Callback-ID header.
func (c *Client) DoAsync(ctx context.Context, method, path string, body any) (*AsyncActionResponse, error) {
	var out any
	callbackID, err := c.DoJSON(ctx, method, path, body, &out)
	if err != nil {
		return nil, err
	}
	return &AsyncActionResponse{Body: out, CallbackID: callbackID}, nil
}
