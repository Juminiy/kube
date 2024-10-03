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
