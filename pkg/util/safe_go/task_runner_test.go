package safe_go

import "testing"

func TestNewRunner(t *testing.T) {
	t.Log(
		NewRunner(
			getTestTasks(),
			WithErrCancel(),
		).Go().Error(),
	)
}
