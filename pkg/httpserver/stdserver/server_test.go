package stdserver

import "testing"

// could not debug in darwin why?
// +passed windows
func TestListenAndServeInfo(t *testing.T) {
	ListenAndServeInfoF(false, 9090)
}
