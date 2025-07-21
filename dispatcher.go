package gitc

import (
	"fmt"
	"sync"
)

type Dispatcher struct {
	tasks map[string]*Task
	mu    sync.RWMutex
}

// initialize a new task dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		tasks: make(map[string]*Task),
	}
}

// stop the task dispatcher
func (d *Dispatcher) StopTask(name string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	task, exists := d.tasks[name]
	if !exists {
		return fmt.Errorf("task %s not found", name)
	}

	close(task.Mailbox)
	delete(d.tasks, name)
	return nil
}

// send a message to a specific task
func (d *Dispatcher) Send(from, to string, Msgtype MessageType, payload interface{}) error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	receiver, exists := d.tasks[to]
	if !exists {
		return fmt.Errorf("receiver %s not found", to)
	}

	msg := Message{
		From:    from,
		To:      to,
		Payload: payload,
	}

	receiver.Mailbox <- msg
	return nil
}

// Start task initialize a new Task (gitc endpoint) and link a message handler
func (d *Dispatcher) StartTask(name string, handler TaskHandler, mailboxSize int) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.tasks[name]; exists {
		return fmt.Errorf("task %s already exists", name)
	}

	task := &Task{
		Name:    name,
		Mailbox: make(chan Message, mailboxSize),
		handler: handler,
	}
	d.tasks[name] = task

	go func() {
		for msg := range task.Mailbox {
			handler(msg)
		}
	}()

	return nil
}

// Reset clears all tasks from the dispatcher.
func (d *Dispatcher) Reset() {
	d.mu.Lock()
	defer d.mu.Unlock()

	for name, task := range d.tasks {
		close(task.Mailbox)
		delete(d.tasks, name)
	}
}
