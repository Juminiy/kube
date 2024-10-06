package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"net"
	"strings"
)

// LookupIP
//
// +example 192.168.31.254 -> 192.168.31.254
// +example app.local -> 192.168.31.254
// +example 192.168.31.254:80 -> 192.168.31.254:80
// +example app.local:80 -> 192.168.31.254:80
func LookupIP(host string) string {
	ipPort := strings.Split(host, ":")
	ip := ""

	if len(ipPort) > 0 {
		ips, err := net.LookupIP(ipPort[0])
		if err != nil {
			stdlog.ErrorF("net LookupIP host: %s error: %s", host, err.Error())
			return host
		}
		if len(ips) < 1 {
			return host
		}
		ip = ips[0].String()
	}

	if len(ipPort) > 1 {
		return StringConcat(ip, ":", ipPort[1])
	}

	return ip
}

func IPStringFromAddr(addr []net.Addr) []string {
	ip := make([]string, len(addr))
	for addrI, addrE := range addr {
		ip[addrI] = addrE.String()
	}
	return ip
}

func IPFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch addr.(type) {
	case *net.IPAddr:
		ip = addr.(*net.IPAddr).IP
	case *net.IPNet:
		ip = addr.(*net.IPNet).IP
	case *net.TCPAddr:
		ip = addr.(*net.TCPAddr).IP
	case *net.UDPAddr:
		ip = addr.(*net.UDPAddr).IP
	case *net.UnixAddr:
		ip = addr.(*net.IPAddr).IP
	}
	return ip
}

// IsIPv4
// TODO: fixbug
func IsIPv4(addr net.Addr) bool {
	return len(IPFromAddr(addr)) == net.IPv4len
}

// IsIPv6
// TODO: fixbug
func IsIPv6(addr net.Addr) bool {
	return len(IPFromAddr(addr)) == net.IPv6len
}
