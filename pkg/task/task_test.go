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

	time.Sleep(10 * time.Millisecond)

	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusIdle, runner.Status, "runner %s is not idle", runner.ID)
	}

	controller.Stop()

	time.Sleep(10 * time.Millisecond)

	for _, runner := range controller.runners {
		require.Equal(t, RunnerStatusStopped, runner.Status, "runner %s is not stopped", runner.ID)
	}
}
