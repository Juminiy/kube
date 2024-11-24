package safe_validator

var (
	_tag    = "valid"
	_simple = false
)

func SetTag(tag string) {
	_tag = tag
}

func SetSimple() {
	_simple = true
}
