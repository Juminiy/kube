package udp

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"net"
	"time"
)

func IPv4Client(serverAddr string) {
	serverUDPAddr := ParseUDPAddr(serverAddr)
	udpConn, err := net.DialUDP(NetworkUDP, nil, serverUDPAddr)
	util.Must(err)
	defer util.SilentCloseIO("udp connection client", udpConn)

	for {
		time.Sleep(util.TimeSecond(1))
		_, err = udpConn.Write(str2Bs(HeartBeatMsgStr))
		if err != nil {
			util.SilentErrorf("write to udp", err)
			continue
		}

		buf := util.GetBuffer()
		_, err = udpConn.Read(buf)
		if err != nil {
			util.SilentError(err)
			continue
		}

		stdlog.InfoF("read from udp server: %s, content: %s", serverAddr, bs2Str(buf))
		util.PutBuffer(buf)

	}

}
