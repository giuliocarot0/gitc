package gitc

import (
	"sync"
)

// Global dispatcher
var (
	GlobalDispatcher *Dispatcher
	once             sync.Once
)

// Initialize the dispatcher once
func init() {
	once.Do(func() {
		GlobalDispatcher = NewDispatcher()
	})
}

/* wrappers for Dispatcher singleton */
func StartTask(name string, handler TaskHandler, mailboxSize int) error {
	return GlobalDispatcher.StartTask(name, handler, mailboxSize)
}

func Send(from, to string, messageType MessageType, payload interface{}) error {
	return GlobalDispatcher.Send(from, to, messageType, payload)
}

func StopTask(name string) error {
	return GlobalDispatcher.StopTask(name)
}

func ResetDispatcher() {
	GlobalDispatcher.Reset()
	GlobalDispatcher = NewDispatcher()
}
