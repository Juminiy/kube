package safe_cast

import (
	"github.com/spf13/cast"
	"math"
)

func F32tof64(f32 float32) float64 {
	return float64(f32)
}

func F64tof32(f64 float64) float32 {
	if f64 > math.MaxFloat32 { // f64 < math.SmallestNonzeroFloat32
		castOverflowErrorF("float64", "float32", f64)
		return InvalidF32
	}
	return float32(f64)
}

func FtoI64[F ~float32 | float64](v F) int64 {
	return cast.ToInt64(cast.ToString(v))
}

func FtoU64[F ~float32 | float64](v F) uint64 {
	return cast.ToUint64(cast.ToString(v))
}
