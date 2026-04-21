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

		switch logs[0].Status {
		case EventFinished:
			return &WaitForEventResult{Data: logs[0].Message}, nil
		case EventError:
			return nil, fmt.Errorf("%s", logs[0].Message)
		}

		if err := waitForNextPoll(ctx, 3*time.Second); err != nil {
			return nil, err
		}
	}
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
