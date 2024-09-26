package util

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
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
