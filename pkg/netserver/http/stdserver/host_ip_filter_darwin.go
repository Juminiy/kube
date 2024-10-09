//go:build darwin

package stdserver

import (
	"github.com/Juminiy/kube/pkg/util"
	"net"
)

func HostListenIPFilter(addr net.Addr) bool {
	switch {
	case util.IsIPv4(addr):
		return !util.ElemIn(util.GetIPv4Str(addr), "0.0.1.1")
	case util.IsIPv6(addr):
		return !util.StringPrefixIn(util.GetIPv6Str(addr), "fe80")
	default:
		return true
	}
}
