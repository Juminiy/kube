package util

import "strconv"

func I64toa(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}

func U64toa(u64 uint64) string {
	return strconv.FormatUint(u64, 10)
}
