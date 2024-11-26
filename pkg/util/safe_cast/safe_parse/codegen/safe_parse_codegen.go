package main

import (
	"github.com/Juminiy/kube/pkg/util"
	kubefile "github.com/Juminiy/kube/pkg/util/file"
	"strings"
)

func main() {
	wd := util.GetWorkPath("safe_parse.go")
	w := kubefile.CodeGen(wd)

	w.Words("// Package safe_parse codegen by codegen/safe_parse_codegen.go, do not edit.").Line()
	w.Words("package", "safe_parse").Line()
	w.Line()
	w.Words("import", "util", util.StringQuote("github.com/Juminiy/kube/pkg/util")).Line()
	w.Line()

	for _, typ := range []string{
		"bool",
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"float32", "float64",
	} {
		upperTyp := strings.ToUpper(string(typ[0])) + typ[1:]
		w.Words("func").
			Word("Parse", upperTyp, "Ok").
			Words("(", "s", "string", ")").
			Words("(", "v", typ, ",", "ok", "bool", ")", "{").Line()

	}

}
