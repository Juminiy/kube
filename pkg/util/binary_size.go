package util

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
