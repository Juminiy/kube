package zero_reflect

// global config
var (
	_noPointerLevel int
)

const (
	_noPointer = iota + 1
	_structSpec
	_mustComparable
)

func Init() {
	initMetaType()
}
