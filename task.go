package gitc

// TaskHandler processes incoming messages
type TaskHandler func(msg Message)

// Task represents a running task
type Task struct {
	Name    string
	Mailbox chan Message
	handler TaskHandler
}
