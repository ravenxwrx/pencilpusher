package task

import (
	"context"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func setDebugLogging() {
	// Set the logging level to debug
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)
}

func TestController(t *testing.T) {
	// Create a new controller
	controller := NewController()

	require.Equal(t, len(controller.runners), RunnerCount())

	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusUnstarted, runner.Status)
	}

	controller.Start()

	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusIdle, runner.Status, "runner %s is not idle", runner.ID)
	}

	controller.Stop()

	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusStopped, runner.Status, "runner %s is not stopped", runner.ID)
	}
}

type MockTask struct {
	ID        uuid.UUID
	Status    string
	Context   map[string]any
	lifecycle chan Event
}

func (t *MockTask) GetID() uuid.UUID {
	return t.ID
}

func (t *MockTask) GetStatus() string {
	return t.Status
}

func (t *MockTask) GetLifecycle() chan Event {
	return t.lifecycle
}

func (t *MockTask) Run(ctx context.Context) error {
	t.Status = TaskStatusRunning
	t.lifecycle <- EventStart{Context: map[string]any{"task_id": t.ID}}

	time.Sleep(1 * time.Second)
	spew.Dump(t.Context)

	t.Status = TaskStatusCompleted
	t.lifecycle <- EventCompleted{Context: map[string]any{"task_id": t.ID}}

	return nil
}

func TestTaskRun(t *testing.T) {
	runnerCount = 2
	controller := NewController()

	require.Equal(t, runnerCount, len(controller.runners))

	id := uuid.New()
	task := &MockTask{
		ID:        id,
		Status:    TaskStatusPending,
		lifecycle: make(chan Event),
		Context:   map[string]any{"my_task_id": id},
	}

	controller.Start()

	controller.queue <- task

	time.Sleep(15 * time.Millisecond)

	require.Equal(t, TaskStatusRunning, task.GetStatus())
	time.Sleep(1 * time.Second)
	require.Equal(t, TaskStatusCompleted, task.GetStatus())
	controller.Stop()
}
