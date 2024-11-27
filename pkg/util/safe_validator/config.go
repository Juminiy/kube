package safe_validator

var (
	_tag   = "valid"
	_debug = false
)

func SetTag(tag string) {
	_tag = tag
}

type Config struct {
	Tag            string
	OnErrorStop    bool
	IgnoreTagError string
	IndirectValue  bool
}
