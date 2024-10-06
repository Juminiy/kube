package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"net"
)

func netInterface() []net.Interface {
	intfs, err := net.Interfaces()
	if err != nil {
		stdlog.ErrorF("psutil network interface error: %s", err.Error())
		return nil
	}

	return intfs
}
