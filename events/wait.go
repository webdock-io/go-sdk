package events

import (
	"context"
	"fmt"
	"time"
)

func (s *Events) WaitForEventToEnd(ctx context.Context, callbackID string) (*EventDTO, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	for {
		res, err := s.List(ctx, ListEventsOptions{CallbackId: &callbackID})
		if err != nil {
			return nil, err
		}
		if len(res.Events) == 0 {
			if err := waitForNextPoll(ctx, 3*time.Second); err != nil {
				return nil, err
			}
			continue
		}

		event := res.Events[0]
		switch event.Status {
		case "finished":
			return &event, nil
		case "error":
			return nil, fmt.Errorf("%s", event.Message)
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
