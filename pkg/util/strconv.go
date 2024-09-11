package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"strconv"
)

func I64toa(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

func U64toa(u64 uint64) string {
	return strconv.FormatUint(u64, 10)
}

func AtoI64(s string) int64 {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		stdlog.Error(err)
		return 0
	}
	return i64
}

func AtoU64(s string) uint64 {
	u64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		stdlog.Error(err)
		return 0
	}
	return u64
}
