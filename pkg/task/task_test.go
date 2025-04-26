package task

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestController(t *testing.T) {
	// Create a new controller
	controller := NewController()

	require.Equal(t, len(controller.runners), RunnerCount())
	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusUnstarted, runner.Status)
	}

	controller.Start()

	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusIdle, runner.Status)
	}

	stopTest := make(chan struct{})
	timer := time.NewTimer(5 * time.Second)

	go func() {
		for range timer.C {
			stopTest <- struct{}{}
		}
	}()

	<-stopTest
	timer.Stop()
}

// Check if the queue is empty
