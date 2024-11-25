package safe_validator

var (
	_tag       = "valid"
	_simple    = false
	_errorStop = false
)

func SetTag(tag string) {
	_tag = tag
}

func SetSimple() {
	_simple = true
}

func SetErrorStop() {
	_errorStop = true
}
