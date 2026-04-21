package servers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/webdock-io/go-sdk/client"
	"github.com/webdock-io/go-sdk/events"
)

type FetchFileOptions struct {
	ServerSlug string
	FilePath   string
}

type FetchFileResponse struct {
	CallbackID string `json:"callbackId" tfsdk:"callback_id"`
}

func (s *Servers) FetchFileAsync(ctx context.Context, opts FetchFileOptions) (*FetchFileResponse, error) {
	data, err := json.Marshal(map[string]string{"filePath": opts.FilePath})
	if err != nil {
		return nil, fmt.Errorf("marshaling request: %w", err)
	}
	c, err := s.client.Do(ctx, "POST", fmt.Sprintf("v1/servers/%s/fetchFile", opts.ServerSlug), bytes.NewBuffer(data), nil)
	if err != nil {
		return nil, err
	}
	callbackID, _ := c.GetHeader(client.CallbackID)
	return &FetchFileResponse{CallbackID: callbackID}, nil
}

func (s *Servers) FetchFileSync(ctx context.Context, opts FetchFileOptions) (string, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := s.FetchFileAsync(ctx, opts)
	if err != nil {
		return "", fmt.Errorf("failed to initiate file fetch for %q on server %q: %w", opts.FilePath, opts.ServerSlug, err)
	}

	eventsClient := events.New(s.client)
	callbackID := resp.CallbackID

	for {

		result, err := eventsClient.List(ctx, events.ListEventsOptions{
			CallbackId: &callbackID,
		})
		if err != nil {
			return "", fmt.Errorf("failed to retrieve operation status (callback ID: %s): %w", callbackID, err)
		}

		if len(result.Events) > 0 {
			ev := result.Events[0]
			switch ev.Status {
			case "finished":
				return ev.Message, nil
			case "error":
				return "", fmt.Errorf("file fetch failed for %q: %s", opts.FilePath, ev.Message)
			}
		}

		timer := time.NewTimer(3 * time.Second)
		select {
		case <-ctx.Done():
			timer.Stop()
			return "", ctx.Err()
		case <-timer.C:
		}
	}
}

func (s *Servers) FetchFile(ctx context.Context, opts FetchFileOptions) (string, error) {
	return s.FetchFileSync(ctx, opts)
}
