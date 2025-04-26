// global var
package stdserver

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/psutil"
	"log"
	"net"
	"net/http"
	"strconv"
)

// global config
var (
	_host            string
	_port            int
	_readTimeoutSec  int
	_writeTimeoutSec int
	_maxHeaderBytes  int
	_tlsCertFilePath string
	_tlsKeyFilePath  string
)

// global var
var (
	_tls     bool
	_handler http.Handler
	_logger  *log.Logger
	_server  *http.Server
)

func Init() {
	_server = &http.Server{}

	_server.Addr = fmt.Sprintf("%s:%d", _host, _port)
	_server.ReadTimeout = util.TimeSecond(_readTimeoutSec)
	_server.WriteTimeout = util.TimeSecond(_writeTimeoutSec)
	_server.MaxHeaderBytes = _maxHeaderBytes

	if _handler != nil {
		_server.Handler = _handler
	}

	if _logger != nil {
		_server.ErrorLog = _logger
	} else {
		_server.ErrorLog = stdlog.Get()
	}

	ListenAndServeInfoF(_tls, _port)
	var serveErr error
	if _tls {
		serveErr = _server.ListenAndServeTLS(_tlsCertFilePath, _tlsKeyFilePath)
	} else {
		serveErr = _server.ListenAndServe()
	}

	if serveErr != nil {
		stdlog.ErrorF("http server: stdserver serve error: %s", serveErr.Error())
	}

}

func ListenAndServeInfoF(tls bool, port int, addrFilter ...psutil.AddrFilter) {
	appProto := "http"
	if tls {
		appProto = "https"
	}
	ipFuncV := func(addr net.Addr) bool {
		for _, ipfn := range addrFilter {
			if !ipfn(addr) {
				return false
			}
		}
		return HostListenIPFilter(addr)
	}
	ipList := psutil.HostIP(
		func(intf net.Interface) bool {
			return (psutil.RunningInterface(intf) && psutil.UpInterface(intf)) ||
				psutil.LoopbackInterface(intf)
		},
		ipFuncV,
	)

	for _, ip := range ipList {
		if util.IsIPv(ip) == 6 {
			stdlog.InfoF("listen and serve %s://[%s]:%d", appProto, ip, port)
		} else {
			stdlog.InfoF("listen and serve %s://%s:%d", appProto, ip, port)
		}
	}
}

func AllIntfs(port int) string {
	return ":" + strconv.Itoa(port)
}

var LawHTTPMethod = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
}
