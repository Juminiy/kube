package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"reflect"
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

func F64toa(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', 3, 64)
}

func F32toa(f32 float32) string {
	return strconv.FormatFloat(float64(f32), 'f', 3, 32)
}

func AtoF64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		stdlog.Error(err)
		return 0.0
	}
	return f64
}

func AtoF32(s string) float32 {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		stdlog.Error(err)
		return 0.0
	}
	return float32(f64)
}

func Ptr2a(v any) string {
	valOf := reflect.ValueOf(v)
	if valOf.Kind() == reflect.Pointer {
		return strconv.FormatUint(uint64(valOf.Pointer()), 16)
	}
	return ""
}
