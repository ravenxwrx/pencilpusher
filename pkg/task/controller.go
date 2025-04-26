package task

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Controller struct {
	done    chan struct{}
	queue   chan Task
	tasks   map[uuid.UUID]Task
	runners map[uuid.UUID]*Runner
}

func NewController() *Controller {
	runners := make(map[uuid.UUID]*Runner)

	for i := 0; i < RunnerCount(); i++ {
		runner := NewRunner()
		runners[runner.ID] = runner
	}

	return &Controller{
		queue:   make(chan Task),
		tasks:   make(map[uuid.UUID]Task),
		runners: runners,
	}
}

func (c *Controller) Start() error {
	for _, runner := range c.runners {
		go runner.Start(c.queue)
	}

	tt := time.Now()

	err := waitForRunners(c.runners, func(runners map[uuid.UUID]*Runner) bool {
		for _, runner := range runners {
			if runner.Status == RunnerStatusUnstarted {
				return true
			}
		}

		return false
	})

	if err != nil {
		return err
	}

	slog.Debug("All runners started", "time", time.Since(tt).String())

	return nil
}

func (c *Controller) Stop() error {
	close(c.queue)

	err := waitForRunners(c.runners, func(runners map[uuid.UUID]*Runner) bool {
		for _, runner := range runners {
			if runner.Status != RunnerStatusStopped {
				return true
			}
		}

		return false
	})

	return err
}

func (c *Controller) AddTask(task Task) {
	c.queue <- task
	c.tasks[task.GetID()] = task
	slog.Debug("Task added to queue", "task_id", task.GetID())
}

func (c *Controller) GetTask(id uuid.UUID) Task {
	task, ok := c.tasks[id]
	if !ok {
		return nil
	}

	return task
}

var errTimeout = fmt.Errorf("timeout")

func waitForRunners(runners map[uuid.UUID]*Runner, call func(runners map[uuid.UUID]*Runner) bool) error {
	timer := time.NewTicker(100 * time.Millisecond)
	retries := 5

Loop:
	for {
		select {
		case <-timer.C:
			retries--
			if retries <= 0 {
				break Loop
			}

			if call(runners) {
				continue Loop
			}

			break Loop
		}
	}
	timer.Stop()

	if retries <= 0 {
		return errTimeout
	}

	return nil
}
