package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"os"
	"os/signal"
	"syscall"
)

func SigNotify(fn Fn) {
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)
	GoSafe(fn)
	stdlog.FatalF("sys signal received <- %s", <-osSig)
}
