package task

import (
	"log/slog"

	"github.com/google/uuid"
)

type Runner struct {
	ID uuid.UUID

	Status string
}

func NewRunner() *Runner {
	return &Runner{
		ID:     uuid.New(),
		Status: RunnerStatusUnstarted,
	}
}

func (r *Runner) Start(queue chan Task) {
	r.Status = RunnerStatusIdle
	for task := range queue {
		r.Status = RunnerStatusBusy
		done := make(chan struct{})

		go func() {
			for {
				select {
				case event := <-task.GetLifecycle():
					switch event.(type) {
					case EventStart:
						slog.Info("Task started", "task_id", task.GetID())
					case EventFailed:
						slog.Error("Task failed", "task_id", task.GetID())
					case EventCompleted:
						slog.Info("Task completed", "task_id", task.GetID())
					}
				case <-done:
					return
				}
			}
		}()
		task.Run()

		r.Status = RunnerStatusIdle

		close(done)
	}
}
