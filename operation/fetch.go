package operation

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

func (o *Operation) Fetch(ctx context.Context, callbackID string) ([]EventLog, error) {
	u := &url.URL{Path: "/events"}
	q := url.Values{}
	q.Set("callbackId", callbackID)
	u.RawQuery = q.Encode()

	var out []EventLog
	_, err := o.client.Do(ctx, "GET", u.String(), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type WaitForEventResult struct {
	Data string `json:"data" tfsdk:"data"`
}

func (o *Operation) WaitForEventToEnd(ctx context.Context, callbackID string) (*WaitForEventResult, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	for {
		logs, err := o.Fetch(ctx, callbackID)
		if err != nil {
			return nil, err
		}
		if len(logs) == 0 {
			if err := waitForNextPoll(ctx, 3*time.Second); err != nil {
				return nil, err
			}
			continue
		}

		result, done, err := waitResultForEventLogs(logs)
		if err != nil {
			return nil, err
		}
		if done {
			return result, nil
		}

		if err := waitForNextPoll(ctx, 3*time.Second); err != nil {
			return nil, err
		}
	}
}

func waitResultForEventLogs(logs []EventLog) (*WaitForEventResult, bool, error) {
	if len(logs) == 0 {
		return nil, false, nil
	}

	allFinished := true
	message := logs[0].Message

	for _, event := range logs {
		switch event.Status {
		case EventFinished:
			if message == "" {
				message = event.Message
			}
		case EventError:
			return nil, false, fmt.Errorf("%s", event.Message)
		default:
			allFinished = false
		}
	}

	if allFinished {
		return &WaitForEventResult{Data: message}, true, nil
	}

	return nil, false, nil
}

func waitForNextPoll(ctx context.Context, delay time.Duration) error {
	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}
