//go:build linux

package stdserver

import (
	"github.com/Juminiy/kube/pkg/util"
	"net"
)

func HostListenIPFilter(addr net.Addr) bool {
	switch {
	case util.IsIPv4(addr):
		return !util.ElemIn(util.GetIPv4Str(addr),
			"172.17.0.1", "172.18.0.1", "172.20.0.1", //docker0 network
		)
	case util.IsIPv6(addr):
		return !util.StringPrefixIn(util.GetIPv6Str(addr), "fe80")
	default:
		return true
	}
}
