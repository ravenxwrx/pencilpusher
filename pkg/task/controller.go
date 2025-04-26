package task

import (
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Controller struct {
	done    chan struct{}
	queue   chan Task
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
		runners: runners,
	}
}

func (c *Controller) Start() {
	for _, runner := range c.runners {
		go runner.Start(c.queue)
	}

	tt := time.Now()
Loop:
	for {
		for _, runner := range c.runners {
			if runner.Status == RunnerStatusUnstarted {
				continue Loop
			}
		}

		break
	}

	slog.Debug("All runners started", "time", time.Since(tt).String())
}

func (c *Controller) Stop() {
	close(c.queue)

Loop:
	for {
		for _, runner := range c.runners {
			if runner.Status != RunnerStatusStopped {
				continue Loop
			}
		}

		break
	}
}
