package main

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	kubefile "github.com/Juminiy/kube/pkg/util/file"
	kubereflect "github.com/Juminiy/kube/pkg/util/reflect"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"strings"
)

func main() {
	safeCastIntLike()
}

func safeCastIntLike() {
	for fromTypStr, fromV := range intLikeTVMap {
		wd, err := internal_api.GetWorkPath(fromTypStr + ".go")
		util.Must(err)
		// create file
		w := kubefile.CodeGen(wd)

		w.Words("// Package safe_cast codegen by codegen/safe_cast_codegen.go, do not edit.").Line()
		w.Words("package", "safe_cast").Line()
		w.Line()
		w.Words("import", util.StringQuote("math")).Line()
		w.Line()

		fromShort := intLikeShortMap[fromTypStr]

		for toTypStr, toV := range intLikeTVMap {
			toShort := intLikeShortMap[toTypStr]

			// func start
			w.Word("func").
				Word(util.StringConcat(getIntLikeTypeUpper(fromShort), "to", getIntLikeTypeUpper(toShort))).
				Words("(", fromShort, fromTypStr, ")").
				Word(toTypStr).Word("{").Line()

			ifSection := func(comp, rVal, errF string) {
				w.Word("if").Words(fromShort, comp, rVal).Word("{").Line()
				w.Words(errF, "(", util.StringQuote(fromTypStr), ",", util.StringQuote(toTypStr), ",", fromShort, ")").Line()
				w.Words("return", "Invalid"+getIntLikeTypeUpper(toShort)).Line()
				w.Word("}").Line()
			}

			// negative section
			if isI(fromTypStr) && isU(toTypStr) {
				ifSection("<", "0", "castNegativeErrorF")
			}

			// overflow section
			if fromV.Typ.Size() > toV.Typ.Size() && toV.Typ.String() != "uintptr" {
				// positive overflow section
				ifSection(">", "math.Max"+getIntLikeTypeUpper(toTypStr), "castOverflowErrorF")

				// negative overflow section
				if isI(fromTypStr) && isI(toTypStr) {
					ifSection("<", "math.Min"+getIntLikeTypeUpper(toTypStr), "castOverflowErrorF")
				}
			} else if fromV.Typ.Size() == toV.Typ.Size() && isU(fromTypStr) && isI(toTypStr) {
				// positive overflow section: signed bit-lost overflow
				ifSection(">", "math.Max"+getIntLikeTypeUpper(toTypStr), "castOverflowErrorF")
			}

			// return section
			w.Words("return", util.StringConcat(toTypStr, "(", fromShort, ")")).Line()

			// end func
			w.Word("}").Line().Line()
		}
		w.Done()
	}
}

// const: must not be assigned
var (
	intLikeTVMap = map[string]kubereflect.TypVal{
		"int":     kubereflect.IndirectOf(safe_cast.InvalidI),
		"int8":    kubereflect.IndirectOf(safe_cast.InvalidI8),
		"int16":   kubereflect.IndirectOf(safe_cast.InvalidI16),
		"int32":   kubereflect.IndirectOf(safe_cast.InvalidI32),
		"int64":   kubereflect.IndirectOf(safe_cast.InvalidI64),
		"uint":    kubereflect.IndirectOf(safe_cast.InvalidU),
		"uint8":   kubereflect.IndirectOf(safe_cast.InvalidU8),
		"uint16":  kubereflect.IndirectOf(safe_cast.InvalidU16),
		"uint32":  kubereflect.IndirectOf(safe_cast.InvalidU32),
		"uint64":  kubereflect.IndirectOf(safe_cast.InvalidU64),
		"uintptr": kubereflect.IndirectOf(safe_cast.InvalidUPtr),
	}
	intLikeShortMap = map[string]string{
		"int":     "i",
		"int8":    "i8",
		"int16":   "i16",
		"int32":   "i32",
		"int64":   "i64",
		"uint":    "u",
		"uint8":   "u8",
		"uint16":  "u16",
		"uint32":  "u32",
		"uint64":  "u64",
		"uintptr": "uptr",
	}
)

func getIntLikeTypeUpper(typ string) string {
	if typ == "uptr" {
		return "UPtr"
	}

	return strings.ToUpper(string(typ[0])) + typ[1:]
}

func isU(typ string) bool {
	return typ[0] == 'u'
}

func isI(typ string) bool {
	return typ[0] == 'i'
}
