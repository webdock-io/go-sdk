package events

import "testing"

func TestWaitResultForEventsWaitsForAllEvents(t *testing.T) {
	result, done, err := waitResultForEvents([]EventDTO{
		{Status: "finished", Message: "phase one done"},
		{Status: "working", Message: "still provisioning"},
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if done {
		t.Fatal("expected wait to continue while any event is still working")
	}
	if result != nil {
		t.Fatalf("expected no result while waiting, got %#v", result)
	}
}

func TestWaitResultForEventsCompletesWhenAllEventsFinish(t *testing.T) {
	result, done, err := waitResultForEvents([]EventDTO{
		{Status: "finished", Message: "phase one done"},
		{Status: "finished", Message: "phase two done"},
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !done {
		t.Fatal("expected wait to complete when all events are finished")
	}
	if result == nil || result.Message != "phase one done" {
		t.Fatalf("expected first event, got %#v", result)
	}
}

func TestWaitResultForEventsReturnsEventError(t *testing.T) {
	_, done, err := waitResultForEvents([]EventDTO{
		{Status: "finished", Message: "phase one done"},
		{Status: "error", Message: "provisioning failed"},
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if done {
		t.Fatal("expected done to be false when returning an error")
	}
}
