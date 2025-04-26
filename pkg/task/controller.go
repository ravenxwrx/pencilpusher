package task

import (
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
		done:    make(chan struct{}),
		queue:   make(chan Task),
		runners: runners,
	}
}

func (c *Controller) Start() {
	for _, runner := range c.runners {
		go runner.Start(c.queue)
	}

	for {
		done := true

		for _, runner := range c.runners {
			if runner.Status == RunnerStatusUnstarted {
				done = false
			}
		}

		if done {
			break
		}
	}
}
