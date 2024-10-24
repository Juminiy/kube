package util

func NewBool(bVar bool) *bool {
	return &bVar
}

func NewInt32(i32 int32) *int32 {
	return &i32
}

func NewInt64(i64 int64) *int64 {
	return &i64
}

func NewFloat32(f32 float32) *float32 {
	return &f32
}

func NewFloat64(f64 float64) *float64 {
	return &f64
}

func NewString(str string) *string {
	return &str
}

func New[T any](val T) *T {
	return &val
}

func New2[T any](val T) **T { return New(New(val)) }

func NewZero[T any](val T) *T {
	return new(T)
}

func ToElemPtrSlice[ElemSlice []E, ElemPtrSlice []*E, E any](s ElemSlice) ElemPtrSlice {
	eps := make(ElemPtrSlice, len(s))
	for i := range s {
		eps[i] = New(s[i])
	}
	return eps
}

func ToElemPtrMap[ElemMap map[K]E, ElemPtrMap map[K]*E, K comparable, E any](m ElemMap) ElemPtrMap {
	epm := make(ElemPtrMap, len(m))
	for k, v := range m {
		epm[k] = New(v)
	}
	return epm
}
