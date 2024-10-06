package stdserver

import "testing"

// could not debug in darwin why?
func TestListenAndServeInfo(t *testing.T) {
	ListenAndServeInfo(false, 9090)
}
