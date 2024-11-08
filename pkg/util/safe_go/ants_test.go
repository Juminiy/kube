package safe_go

import "testing"

func TestAntsRunner(t *testing.T) {
	AntsInit()
	testRunT(t, AntsRunner)
}
