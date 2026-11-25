package servers

import (
	"context"
	"fmt"

	"github.com/webdock-io/go-sdk/operation"
)

type FetchFileOptions struct {
	ServerSlug string
	FilePath   string
}

type FetchFileResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

// StartFileFetch queues a file fetch and returns its callback ID. Mirrors
// servers.StartFileFetch() in the Node SDK.
func (s *Servers) StartFileFetch(ctx context.Context, opts FetchFileOptions) (*FetchFileResponse, error) {
	callbackID, err := s.client.DoJSON(ctx, "POST", fmt.Sprintf("/servers/%s/fetchFile", opts.ServerSlug), map[string]string{"filePath": opts.FilePath}, nil)
	if err != nil {
		return nil, err
	}
	return &FetchFileResponse{CallbackID: callbackID}, nil
}

// Deprecated: use StartFileFetch.
func (s *Servers) FetchFileAsync(ctx context.Context, opts FetchFileOptions) (*FetchFileResponse, error) {
	return s.StartFileFetch(ctx, opts)
}

// FetchFile starts a file fetch, waits for it to finish and returns the file
// content as base64. Mirrors servers.fetchFile() in the Node SDK.
func (s *Servers) FetchFile(ctx context.Context, opts FetchFileOptions) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	init, err := s.StartFileFetch(ctx, opts)
	if err != nil {
		return "", fmt.Errorf("failed to initiate file fetch for %q on server %q: %w", opts.FilePath, opts.ServerSlug, err)
	}

	ops := operation.New(s.client)
	result, err := ops.WaitForEventToEnd(ctx, init.CallbackID)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve operation status (callback ID: %s): %w", init.CallbackID, err)
	}
	return result.Data, nil
}

// Deprecated: use FetchFile.
func (s *Servers) FetchFileSync(ctx context.Context, opts FetchFileOptions) (string, error) {
	return s.FetchFile(ctx, opts)
}
