package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	unit "github.com/docker/go-units"
	"github.com/spf13/cast"
	"strconv"
)

// <binarySI>        ::= Ki | Mi | Gi | Ti | Pi | Ei
// refer from: resource/quantity.go
const (
	Ki = 1 << 10
	Mi = 1 << 20
	Gi = 1 << 30
	Ti = 1 << 40
	Pi = 1 << 50
	Ei = 1 << 60
)

// BytesOf
// get a proper measurement: B/KiB/MiB/GiB/TiB/PiB/EiB
func BytesOf(bs []byte) string {
	lenOf := len(bs)
	var appr float64
	var measure = "B"
	if lenOf < Ki {
		appr = float64(lenOf)
	} else if lenOf < Mi {
		appr = BytesKB(bs)
		measure = "KiB"
	} else if lenOf < Gi {
		appr = BytesMB(bs)
		measure = "MiB"
	} else if lenOf < Ti {
		appr = BytesGB(bs)
		measure = "GiB"
	} else {
		stdlog.WarnF("too much byte %d B, do not convert", lenOf)
		return strconv.Itoa(lenOf) + measure
	}
	return F64toa(appr) + measure
}

func BytesKB(bs []byte) float64 {
	return float64(len(bs)) / (1.0 * Ki)
}

func BytesMB(bs []byte) float64 {
	return float64(len(bs)) / (1.0 * Mi)
}

func BytesGB(bs []byte) float64 {
	return float64(len(bs)) / (1.0 * Gi)
}

func MeasureByte(size int) string {
	var (
		appr float64 = 0.0
		meas string  = "B"
	)
	switch {
	case size < 0:
		return ""
	case size < Ki:
		appr = float64(size)
	case size < Mi:
		appr = float64(size) / (1.0 * Ki)
		meas = "KiB"
	case size < Gi:
		appr = float64(size) / (1.0 * Mi)
		meas = "MiB"
	case size < Ti:
		appr = float64(size) / (1.0 * Gi)
		meas = "GiB"
	case size < Pi:
		appr = float64(size) / (1.0 * Ti)
		meas = "TiB"
	case size < Ei:
		appr = float64(size) / (1.0 * Pi)
		meas = "Pi"
	default:
		stdlog.WarnF("too much byte %d B, do not convert", size)
		return strconv.Itoa(size) + meas
	}
	return F64toa(appr) + meas
}

// +parameter a: int, uint, float
func HumanSize(a any) string {
	return unit.HumanSize(cast.ToFloat64(a))
}

// +parameter a: int, uint, float
func HumanBytesSize(a any) string {
	return unit.BytesSize(cast.ToFloat64(a))
}
