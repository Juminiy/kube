package util

import "strconv"

func I64toa(i64 int64) string {
	return strconv.FormatInt(i64, 10)
}
