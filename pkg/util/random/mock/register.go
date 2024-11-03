package mock

import "github.com/Juminiy/kube/pkg/util"

// goroutine is unsafe
// registerFunc is not concurrent safe, please register as project init, do not dynamic set and get

func RegisterBool(name string, fn BoolFunc) bool {
	return register(boolFunc, name, fn)
}

func RegisterUint(name string, fn UintFunc) bool {
	return register(uintFunc, name, fn)
}

func RegisterInt(name string, fn IntFunc) bool {
	return register(intFunc, name, fn)
}

func RegisterFloat(name string, fn FloatFunc) bool {
	return register(floatFunc, name, fn)
}

func RegisterString(name string, fn StringFunc) bool {
	return register(stringFunc, name, fn)
}

func register[Map ~map[string]V, V any](m Map, name string, fn V) bool {
	if util.MapOk(m, name) {
		return false
	}
	m[name] = fn
	return true
}
