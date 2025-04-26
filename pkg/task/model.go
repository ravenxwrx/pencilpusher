package task

import (
	"github.com/google/uuid"
)

const (
	TaskStatusPending   = "pending"
	TaskStatusRunning   = "running"
	TaskStatusCompleted = "completed"
	TaskStatusFailed    = "failed"

	RunnerStatusUnstarted = "not_started"
	RunnerStatusIdle      = "idle"
	RunnerStatusBusy      = "busy"
)

type Event interface {
	GetContext() map[string]any
}

type EventStart struct {
	Context map[string]any
}

func (e EventStart) GetContext() map[string]any {
	return e.Context
}

type EventFailed struct {
	Context map[string]any
}

func (e EventFailed) GetContext() map[string]any {
	return e.Context
}

type EventCompleted struct {
	Context map[string]any
}

func (e EventCompleted) GetContext() map[string]any {
	return e.Context
}

type Task interface {
	GetID() uuid.UUID
	GetStatus() string
	GetLifecycle() chan Event
	Run() error
}
