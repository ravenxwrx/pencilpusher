package task

var (
	runnerCount = 8
)

func RunnerCount() int {
	return runnerCount
}

func SetRunnerCount(count int) {
	if count < 1 {
		panic("runner count must be greater than 0")
	}
	runnerCount = count
}
