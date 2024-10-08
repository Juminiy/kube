package udp

import (
	"bufio"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"net"
	"time"
)

func IPv4Client(serverAddr string) {
	for {
		serverUDPAddr := ParseUDPAddr(serverAddr)
		udpConn, err := net.DialUDP(NetworkUDP, nil, serverUDPAddr)
		util.Must(err)

		time.Sleep(util.TimeSecond(1))
		_, err = udpConn.Write(str2Bs(HeartBeatMsgStr))
		if err != nil {
			util.SilentErrorf("write to udp", err)
			continue
		}

		content, err := bufio.NewReader(udpConn).ReadString('\n')
		if err != nil {
			util.SilentError(err)
			continue
		}

		stdlog.InfoF("read from udp server: %s, content: %s", serverAddr, content)
		util.SilentCloseIO("udp connection client", udpConn)
	}

}
