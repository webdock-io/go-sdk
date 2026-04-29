package operation

import "testing"

func TestWaitResultForEventLogsWaitsForAllEvents(t *testing.T) {
	result, done, err := waitResultForEventLogs([]EventLog{
		{Status: EventFinished, Message: "phase one done"},
		{Status: EventWorking, Message: "still provisioning"},
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

func TestWaitResultForEventLogsCompletesWhenAllEventsFinish(t *testing.T) {
	result, done, err := waitResultForEventLogs([]EventLog{
		{Status: EventFinished, Message: "phase one done"},
		{Status: EventFinished, Message: "phase two done"},
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !done {
		t.Fatal("expected wait to complete when all events are finished")
	}
	if result == nil || result.Data != "phase one done" {
		t.Fatalf("expected first event message, got %#v", result)
	}
}

func TestWaitResultForEventLogsReturnsEventError(t *testing.T) {
	_, done, err := waitResultForEventLogs([]EventLog{
		{Status: EventFinished, Message: "phase one done"},
		{Status: EventError, Message: "provisioning failed"},
	})
	if err == nil {
		t.Fatal("expected error")
	}
	if done {
		t.Fatal("expected done to be false when returning an error")
	}
}
