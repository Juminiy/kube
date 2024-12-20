package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

type ExportStruct struct {
	Int    int
	Float  float64
	String string
}

var esvS = []any{
	ExportStruct{},               // Struct
	&ExportStruct{},              // Pointer to Struct
	util.New(&ExportStruct{}),    // Pointer to Pointer to Struct
	nil,                          // nil
	util.Zero[*ExportStruct](),   // nil Pointer to Struct
	struct{ I int }{},            // unnamed Struct
	&struct{ I int }{},           // Pointer to unnamed Struct
	util.New(&struct{ I int }{}), // Pointer to Pointer to unnamed Struct
	[1]ExportStruct{},            // Array Struct
	&[1]ExportStruct{},           // Pointer to Array Struct
	&[1]*ExportStruct{},          // Pointer to Array Pointer to Struct
	[]ExportStruct{},             // Slice Struct
	&[]ExportStruct{},            // Pointer to Slice Struct
	&[]*ExportStruct{},           // Pointer to Slice Pointer to Struct
}

func TestIndirect(t *testing.T) {
	for _, e := range esvS {
		ei := Indirect(e)
		t.Logf("vi: %v, type: %v", ei.V, ei.T)
	}
}
