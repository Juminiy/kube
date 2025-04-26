package safe_cast

import (
	"math"
)

func F32tof64(f32 float32) float64 {
	return float64(f32)
}

func F64tof32(f64 float64) float32 {
	if f64 < -math.MaxFloat32 || f64 > math.MaxFloat32 { // f64 < math.SmallestNonzeroFloat32
		castOverflowErrorF("float64", "float32", f64)
		return InvalidF32
	}
	return float32(f64)
}
