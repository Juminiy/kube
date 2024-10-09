//go:build linux

package stdserver

func HostListenIPFilter(addr net.Addr) bool {
	return true
}
