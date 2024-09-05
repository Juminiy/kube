package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"strconv"
	"time"
)

func CST2Timestamp(cst string) string {
	cstInt64, err := strconv.ParseInt(cst, 10, 64)
	if err != nil {
		stdlog.ErrorF("cst string: %s error: %s", cst, err.Error())
		return ""
	}
	return time.
		Unix(cstInt64, 0).
		Format(time.DateTime)
}

func Timestamp2CST(timestamp string) string {
	tm, err := time.Parse(time.DateTime, timestamp)
	if err != nil {
		stdlog.ErrorF("timestamp: %s parse to CST time error: %s", timestamp, err.Error())
		return ""
	}
	return strconv.Itoa(int(tm.Unix()))
}
