package zero_reflect

var (
	_metaValue = []any{
		false,
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		float32(0),
		float64(0),
		complex(float32(0), float32(0)),
		complex(float64(0), float64(0)),
		"",
		int(0),
		uint(0),
		uintptr(0),
		struct{}{},
	}
)

func initMetaType() {
	for i := range _metaValue {
		TypeOf(_metaValue[i])
	}
}
