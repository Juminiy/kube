package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"net"
)

func HostIP(intfFilter InterfaceFilter, addrFilter AddrFilter) []string {
	return hostIP("unknown", intfFilter, addrFilter)
}

func HostMAC() []string {
	intfs := netInterface()
	macAddr := make([]string, len(intfs))
	for intfI, intf := range intfs {
		macAddr[intfI] = util.Bytes2StringNoCopy(intf.HardwareAddr)
	}
	return macAddr
}

func HostIPv4() []string {
	return hostIP(
		"ipv4",
		nil,
		util.IsIPv4,
	)
}

func HostIPv6() []string {
	return hostIP(
		"ipv6",
		nil,
		util.IsIPv6,
	)
}

type InterfaceFilter func(net.Interface) bool

func allInterface(net.Interface) bool {
	return true
}

func UpInterface(intf net.Interface) bool {
	return intf.Flags&net.FlagUp != 0
}

func BroadcastInterface(intf net.Interface) bool {
	return intf.Flags&net.FlagBroadcast != 0
}

func LoopbackInterface(intf net.Interface) bool {
	return intf.Flags&net.FlagLoopback != 0
}

func PointToPointInterface(intf net.Interface) bool {
	return intf.Flags&net.FlagPointToPoint != 0
}

func MulticastInterface(intf net.Interface) bool {
	return intf.Flags&net.FlagMulticast != 0
}

func RunningInterface(intf net.Interface) bool {
	return intf.Flags&net.FlagRunning != 0
}

type AddrFilter func(net.Addr) bool

func AllAddr(addr net.Addr) bool {
	return allAddr(addr)
}

func allAddr(net.Addr) bool {
	return true
}

func hostIP(
	area string,
	intfFilter InterfaceFilter,
	addrFilter AddrFilter,
) []string {
	intfs := netInterface()
	ip := make([]string, 0, len(intfs)<<1)
	if intfFilter == nil {
		intfFilter = allInterface
	}
	if addrFilter == nil {
		addrFilter = allAddr
	}
	for _, intf := range intfs {
		if intfFilter(intf) {
			addrs, err := intf.Addrs()
			if err != nil {
				stdlog.ErrorF("psutil host %s address error: %s", area, err.Error())
				return nil
			}
			ip = append(ip, ipFromAddr(addrs, addrFilter)...)
		}
	}
	return ip
}

func netInterface() []net.Interface {
	intfs, err := net.Interfaces()
	if err != nil {
		stdlog.ErrorF("psutil network interface error: %s", err.Error())
		return nil
	}

	return intfs
}

func ipFromAddr(addr []net.Addr, filter AddrFilter) []string {
	ip := make([]string, 0, len(addr))
	for _, addrE := range addr {
		if filter(addrE) {
			ip = append(ip, util.TrimNetMask(addrE.String()))
		}
	}
	return ip
}
