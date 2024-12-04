package fiberserver

import (
	"testing"
)

func TestNew(t *testing.T) {
	New().
		WithPort(8081).
		WithBufferSize("128Mi", "128Mi", "128Mi").
		WithTimeout(60, 60, 60).
		WithName("fiber-web-app").
		WithRESTAPI(DefaultRESTAPI).
		Load()
}
