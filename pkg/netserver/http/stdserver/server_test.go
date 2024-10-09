package stdserver

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

// +passed windows darwin
func TestListenAndServeInfo(t *testing.T) {
	ListenAndServeInfoF(false, 9090, util.IsIPv6)

	ListenAndServeInfoF(false, 9090, util.IsIPv4)
}
