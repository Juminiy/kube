package udp

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"net"
	"time"
)

func IPv4Server(addr string) {
	udpConn, err := net.ListenUDP(NetworkUDP, ParseUDPAddr(addr))
	util.Must(err)
	defer util.SilentCloseIO("udp connection server", udpConn)

	for {
		time.Sleep(util.TimeSecond(1))
		buf := make([]byte, 0, util.MagicBufferCap)
		_, clientAddr, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			util.SilentErrorf("read from udp", err)
			continue
		}

		if len(buf) > 0 {
			stdlog.InfoF("read from client address: %s content: %s", clientAddr.String(), bs2Str(buf))
		}

		_, err = udpConn.WriteToUDP(str2Bs(HeartBeatMsgStr), clientAddr)
		if err != nil {
			util.SilentErrorf("write to udp", err)
			continue
		}
	}

}
