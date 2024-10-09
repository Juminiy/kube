//go:build windows

package stdserver

import "net"

func HostListenIPFilter(addr net.Addr) bool {
	return true
}
