package udp

import (
	"github.com/Juminiy/kube/pkg/util"
	"net"
	"net/netip"
)

func ParseUDPAddr(addr string) *net.UDPAddr {
	return net.UDPAddrFromAddrPort(netip.MustParseAddrPort(addr))
}

const (
	HeartBeatMsgStr = "Ciallo~"
	NetworkUDP      = "udp"
	NetworkUDP4     = "udp4"
	NetworkUDP6     = "udp6"
)

var (
	bs2Str = util.Bytes2StringNoCopy
	str2Bs = util.String2BytesNoCopy
)
