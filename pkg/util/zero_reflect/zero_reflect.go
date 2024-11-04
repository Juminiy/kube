package zero_reflect

// global config
var (
	_noPointerLevel int
)

const (
	_invalid = iota
	_direct
	_noPointer
	_structSpec
	_mustComparable
)

func Init() {
	initMetaType()
}
