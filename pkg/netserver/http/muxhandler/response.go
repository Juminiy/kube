package muxhandler

import (
	"encoding/json"
	"github.com/Juminiy/kube/pkg/util"
	"net/http"
)

// Status
// must http.StatusCode
func Status(resp http.ResponseWriter, code int) {
	resp.WriteHeader(code)
}

func String(resp http.ResponseWriter, v string) {
	_, err := resp.Write(str2Bs(v))
	if err != nil {
		Status(resp, http.StatusInternalServerError)
		return
	}
}

func JSON(resp http.ResponseWriter, v any) {
	bs, err := json.Marshal(v)
	if err != nil {
		Status(resp, http.StatusInternalServerError)
		return
	}
	_, err = resp.Write(bs)
	if err != nil {
		Status(resp, http.StatusInternalServerError)
		return
	}
}

var (
	bs2Str = util.Bytes2StringNoCopy
	str2Bs = util.String2BytesNoCopy
)
