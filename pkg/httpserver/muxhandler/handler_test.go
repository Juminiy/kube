package muxhandler

import (
	"github.com/Juminiy/kube/pkg/httpserver/stdserver"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/psutil"
	"net/http"
	"testing"
)

func TestMuxHandler(t *testing.T) {
	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/api/v1/sys", func(respW http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			respW.WriteHeader(http.StatusNotFound)
			return
		}
		_, err := respW.Write(psutil.Marshal())
		if err != nil {
			respW.WriteHeader(http.StatusInternalServerError)
			return
		}
		respW.WriteHeader(http.StatusOK)
	})

	stdserver.New().
		WithHandler(muxHandler).
		WithHost("0.0.0.0").WithPort(8080).
		WithMaxHeaderBytes(16*util.Mi).
		WithTimeoutSec(10, 10).
		Load()
}
