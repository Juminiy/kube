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

func IsIPv4(addr net.Addr) bool {
	ipTo4 := IPFromAddr(addr).To4()
	return ipTo4 != nil &&
		(len(ipTo4) == net.IPv4len ||
			len(ipTo4) == net.IPv6len && ElemsIn(ipTo4[:12], []byte{0})) &&
		IsIPv(ipTo4.String()) == 4 // most obvious character
}

func IsIPv6(addr net.Addr) bool {
	ipTo16 := IPFromAddr(addr).To16()
	return ipTo16 != nil &&
		//len(ipTo16) == net.IPv6len && !ElemsIn(ipTo16[:12], []byte{0}) && //local mismatch
		IsIPv(ipTo16.String()) == 6 // most obvious character
}

func IsIPv(ip string) uint8 {
	if strings.Contains(ip, ":") {
		return 6
	}
	return 4
}

func GetIPv4Str(addr net.Addr) string {
	return IPFromAddr(addr).To4().String()
}

func GetIPv6Str(addr net.Addr) string {
	return IPFromAddr(addr).To16().String()
}

func TrimNetMask(ip string) string {
	ipMask := strings.Split(ip, "/")
	switch len(ipMask) {
	default:
		return ""
	case 1:
		return ip
	case 2:
		return ipMask[0]
	}
}

func TrimProto(addr string) string {
	dSlashIndex := strings.Index(addr, "://")
	if dSlashIndex != -1 {
		return addr[dSlashIndex+3:]
	}
	return addr
}
