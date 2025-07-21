package tests

import (
	"sync"
	"testing"

	"github.com/giuliocarot0/gitc"
)

// Test basic message sending
func TestSendMessage(t *testing.T) {
	// Reset dispatcher for a clean state
	gitc.ResetDispatcher()

	var wg sync.WaitGroup
	wg.Add(1)

	// Start receiver task
	err := gitc.StartTask("receiver", func(msg gitc.Message) {
		if msg.From != "sender" {
			t.Errorf("unexpected sender: got %q, want %q", msg.From, "sender")
		}
		if msg.Payload != "Hello" {
			t.Errorf("unexpected payload: got %v, want %v", msg.Payload, "Hello")
		}
		wg.Done() // Signal that we received the message
	}, 5)
	if err != nil {
		t.Fatalf("failed to start receiver: %v", err)
	}

	// Start sender task (dummy handler)
	err = gitc.StartTask("sender", func(msg gitc.Message) {}, 5)
	if err != nil {
		t.Fatalf("failed to start sender: %v", err)
	}

	// Send a message
	err = gitc.Send("sender", "receiver", gitc.MSG0, "Hello")
	if err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	// Wait for the receiver to process
	wg.Wait()
}

// Test sending to a non-existent task
func TestSendToNonExistentTask(t *testing.T) {
	gitc.ResetDispatcher()

	err := gitc.Send("anyone", "ghost", gitc.MSG0, "Hello")
	if err == nil {
		t.Error("expected error when sending to non-existent task, got nil")
	}
}

// Test duplicate task names
func TestDuplicateTaskNames(t *testing.T) {
	gitc.ResetDispatcher()

	handler := func(msg gitc.Message) {}
	err := gitc.StartTask("task1", handler, 5)
	if err != nil {
		t.Fatalf("failed to start task1: %v", err)
	}

	// Try starting another task with the same name
	err = gitc.StartTask("task1", handler, 5)
	if err == nil {
		t.Error("expected error for duplicate task name, got nil")
	}
}

// Test ResetDispatcher clears all tasks
func TestResetDispatcher(t *testing.T) {
	gitc.ResetDispatcher()

	handler := func(msg gitc.Message) {}
	_ = gitc.StartTask("taskA", handler, 5)
	_ = gitc.StartTask("taskB", handler, 5)

	// Reset
	gitc.ResetDispatcher()

	// Sending after reset should fail
	err := gitc.Send("taskA", "taskB", gitc.MSG0, "Hello")
	if err == nil {
		t.Error("expected error after ResetDispatcher, got nil")
	}
}
