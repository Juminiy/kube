package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"time"
)

const (
	TimeLocationAsiaShanghai = "Asia/Shanghai"
)

// CST2Timestamp
// Convert (CST: UTC+8) to (Unix timestamp)
func CST2Timestamp(cst string) string {
	cstTm, err := time.ParseInLocation(time.DateTime, cst, cstLocation())
	if err != nil {
		stdlog.ErrorF("convert cst: %s to timestamp error: %s", cst, err.Error())
		return ""
	}
	return I64toa(cstTm.Unix())
}

// Timestamp2CST
// Convert (Unix timestamp) to (CST: UTC+8)
func Timestamp2CST(timestamp string) string {
	return time.Unix(AtoI64(timestamp), 0).
		In(cstLocation()).
		Format(time.DateTime)
}

func cstLocation() *time.Location {
	cstLoc, err := time.LoadLocation(TimeLocationAsiaShanghai)
	if err != nil {
		stdlog.ErrorF("LoadLocation: %s, error: %s", TimeLocationAsiaShanghai, err.Error())
		return nil
	}
	return cstLoc
}
