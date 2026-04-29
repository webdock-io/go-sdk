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

		event, done, err := waitResultForEvents(res.Events)
		if err != nil {
			return nil, err
		}
		if done {
			return event, nil
		}

		if err := waitForNextPoll(ctx, 3*time.Second); err != nil {
			return nil, err
		}
	}
}

func waitResultForEvents(events []EventDTO) (*EventDTO, bool, error) {
	if len(events) == 0 {
		return nil, false, nil
	}

	allFinished := true
	finished := events[0]

	for _, event := range events {
		switch event.Status {
		case "finished":
			if finished.Message == "" {
				finished = event
			}
		case "error":
			return nil, false, fmt.Errorf("%s", event.Message)
		default:
			allFinished = false
		}
	}

	if allFinished {
		return &finished, true, nil
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
