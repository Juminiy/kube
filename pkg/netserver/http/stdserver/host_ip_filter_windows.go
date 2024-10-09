//go:build windows

package stdserver

import (
	"github.com/Juminiy/kube/pkg/util"
	"net"
)

func HostListenIPFilter(addr net.Addr) bool {
	switch {
	case util.IsIPv4(addr):
		return !util.ElemIn(util.GetIPv4Str(addr), "2.0.0.1")
	default:
		return true
	}
}
