package codegen

import "github.com/Juminiy/kube/pkg/util"

var (
	goWordMap = map[string]struct{}{
		"break":       util.NilStruct(),
		"default":     util.NilStruct(),
		"func":        util.NilStruct(),
		"interface":   util.NilStruct(),
		"select":      util.NilStruct(),
		"case":        util.NilStruct(),
		"defer":       util.NilStruct(),
		"go":          util.NilStruct(),
		"map":         util.NilStruct(),
		"struct":      util.NilStruct(),
		"chan":        util.NilStruct(),
		"else":        util.NilStruct(),
		"goto":        util.NilStruct(),
		"package":     util.NilStruct(),
		"switch":      util.NilStruct(),
		"const":       util.NilStruct(),
		"fallthrough": util.NilStruct(),
		"if":          util.NilStruct(),
		"range":       util.NilStruct(),
		"type":        util.NilStruct(),
		"continue":    util.NilStruct(),
		"for":         util.NilStruct(),
		"import":      util.NilStruct(),
		"return":      util.NilStruct(),
		"var":         util.NilStruct(),
		"true":        util.NilStruct(),
		"false":       util.NilStruct(),
		"iota":        util.NilStruct(),
		"nil":         util.NilStruct(),
	}
)
