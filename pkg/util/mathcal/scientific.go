package mathcal

import "math"

func Ln(x float64) float64 {
	return math.Log(x)
}

func LogLog(x float64, l1, l2 int) float64 {
	return LoopLog(x, l1, l2)
}

func LoopLog(x float64, l int, lx ...int) float64 {
	lgx := Ln(x) / Ln(float64(l))
	for _, l := range lx {
		lgx = Ln(lgx) / Ln(float64(l))
	}
	return lgx
}
