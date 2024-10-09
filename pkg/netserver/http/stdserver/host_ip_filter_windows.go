//go:build windows

package stdserver

func HostListenIPFilter(addr net.Addr) bool {
	return true
}
