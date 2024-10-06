package muxhandler

import (
	"github.com/Juminiy/kube/pkg/util/psutil"
	"net/http"
	"testing"
)

func TestMuxHandler(t *testing.T) {
	muxHandler := http.NewServeMux()
	muxHandler.HandleFunc("/api/v1/sys", func(respW http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			respW.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := respW.Write(psutil.Marshal())
		if err != nil {
			respW.WriteHeader(http.StatusInternalServerError)
			return
		}
		respW.WriteHeader(http.StatusOK)
	})
}
