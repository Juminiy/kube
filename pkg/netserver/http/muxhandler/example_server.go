package muxhandler

import (
	"github.com/Juminiy/kube/pkg/netserver/http/stdserver"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/psutil"
	"net/http"
)

func GetHostInfo(resp http.ResponseWriter, req *http.Request) {
	String(resp, util.Bytes2StringNoCopy(psutil.Marshal()))
}

func Pong(resp http.ResponseWriter, req *http.Request) {
	String(resp, "pong")
}

func ExampleHTTPServerRun() {
	stdserver.New().
		WithHandler(New(map[string]HandlerV{
			"/api/v1/host": GET(GetHostInfo),
			"/api/v1/ping": ANY(Pong),
		})).
		WithHost("0.0.0.0").WithPort(8080).
		WithMaxHeaderBytes(16*util.Mi).
		WithTimeoutSec(10, 10).
		Load()
}
