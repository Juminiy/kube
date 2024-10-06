package util

import "github.com/Juminiy/kube/pkg/log_api/stdlog"

func castErrorF(fromTyp, toTyp string, v any, err string) {
	stdlog.ErrorF("cast %s(%v) to %s, error: %s", fromTyp, v, toTyp, err)
}
