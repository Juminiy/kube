package util

import (
	"testing"
)

// +passed
func TestTimestamp2CST(t *testing.T) {
	t.Log(Timestamp2CST("1136185445"))
}

// +passed
func TestCST2Timestamp(t *testing.T) {
	t.Log(CST2Timestamp("2006-01-02 15:04:05"))
}
