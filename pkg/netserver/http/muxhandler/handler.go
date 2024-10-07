package muxhandler

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"net/http"
	"strings"
	"time"
)

type HandlerMap map[string]HandlerV

type HandlerV struct {
	handlerFunc http.HandlerFunc
	method      uint16
}

func GET(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodGet)
}

func HEAD(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodHead)
}

func POST(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodPost)
}

func PUT(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodPut)
}

func PATCH(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodPatch)
}

func DELETE(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodDelete)
}

func CONNECT(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodConnect)
}

func OPTIONS(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodOptions)
}

func TRACE(handlerFunc http.HandlerFunc) HandlerV {
	return handle(handlerFunc, http.MethodTrace)
}

func ANY(handlerFunc http.HandlerFunc) HandlerV {
	return handle(
		handlerFunc,
		http.MethodGet, http.MethodHead, http.MethodPost,
		http.MethodPut, http.MethodPatch, http.MethodDelete,
		http.MethodConnect, http.MethodOptions, http.MethodTrace,
	)
}

func Handle(handlerFunc http.HandlerFunc, httpMethod ...string) HandlerV {
	return handle(handlerFunc, httpMethod...)
}

func handle(handlerFunc http.HandlerFunc, httpMethod ...string) HandlerV {
	handlerV := HandlerV{
		handlerFunc: handlerFunc,
	}
	for _, method := range httpMethod {
		handlerV.method |= methodMap[strings.ToUpper(method)]
	}
	return handlerV
}

// referred from: net/http/method.go
var methodMap = map[string]uint16{
	http.MethodGet:     1 << 0,
	http.MethodHead:    1 << 1,
	http.MethodPost:    1 << 2,
	http.MethodPut:     1 << 3,
	http.MethodPatch:   1 << 4,
	http.MethodDelete:  1 << 5,
	http.MethodConnect: 1 << 6,
	http.MethodOptions: 1 << 7,
	http.MethodTrace:   1 << 8,
}

func (h *HandlerV) allow() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		fromTime := time.Now()
		if methodMap[req.Method]&h.method == 0 {
			resp.WriteHeader(http.StatusNotFound)
			return
		}
		h.handlerFunc(resp, req)
		h.accessLog(fromTime)(resp, req)
	}
}

func (h *HandlerV) accessLog(fromTime time.Time) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		stdlog.InfoF("%s %s %s %s", req.RemoteAddr, req.Method, req.RequestURI, util.MeasureTime(time.Since(fromTime)))
	}
}

func New(hm HandlerMap) *http.ServeMux {
	muxHandler := http.NewServeMux()

	for pattern, handlerV := range hm {
		muxHandler.HandleFunc(pattern, handlerV.allow())
	}

	return muxHandler
}
