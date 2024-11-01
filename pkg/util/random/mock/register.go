package mock

func RegisterBool(name string, fn BoolFunc) bool {
	return true
}

func RegisterUint(name string, fn UintFunc) bool {
	return true
}

func RegisterInt(name string, fn IntFunc) bool {
	return true
}

func RegisterFloat(name string, fn FloatFunc) bool {
	return true
}

func RegisterString(name string, fn StringFunc) bool {
	return true
}
