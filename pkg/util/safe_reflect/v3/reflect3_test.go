package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

type ExportStruct struct {
	Int    int     `gorm:"column:c_int;type:int;comment:int_ref" json:"int,inline"`
	Float  float64 `gorm:"column:c_float;type:float;comment:float_ref" json:"float,inline"`
	String string  `gorm:"column:c_str;type:varchar(5);comment:str_ref" json:"string,inline"`
}

var esvS = map[string]any{
	"Struct":                       ExportStruct{},            // Struct
	"Pointer to Struct":            &ExportStruct{},           // Pointer to Struct
	"Pointer to Pointer to Struct": util.New(&ExportStruct{}), // Pointer to Pointer to Struct
	//"nil":                                     nil,                          // nil
	"nil Pointer to Struct":                   util.Zero[*ExportStruct](),   // nil Pointer to Struct
	"unnamed Struct":                          struct{ I int }{},            // unnamed Struct
	"Pointer to unnamed Struct":               &struct{ I int }{},           // Pointer to unnamed Struct
	"Pointer to Pointer to unnamed Struct":    util.New(&struct{ I int }{}), // Pointer to Pointer to unnamed Struct
	"Array Elem Struct":                       [1]ExportStruct{},            // Array Elem Struct
	"Pointer to Array Elem Struct":            &[1]ExportStruct{},           // Pointer to Array Elem Struct
	"Pointer to Array Elem Pointer to Struct": &[1]*ExportStruct{},          // Pointer to Array Elem Pointer to Struct
	"Slice Elem Struct":                       []ExportStruct{},             // Slice Elem Struct
	"Pointer to Slice Elem Struct":            &[]ExportStruct{},            // Pointer to Slice Elem Struct
	"Pointer to Slice Elem Pointer to Struct": &[]*ExportStruct{},           // Pointer to Slice Elem Pointer to Struct
}

var Pretty = safe_json.Pretty
var String = safe_json.String

func TestIndirect(t *testing.T) {
	for desc, e := range esvS {
		ei := Indirect(e)
		t.Logf("desc: [%40s] type: [%40v] value: [%+40v] ", desc, ei.T, ei.V)
	}
}

func TestIndirect_TypedPointer(t *testing.T) {
	for _, e := range []any{
		util.Zero[*ExportStruct](),
		util.Zero[*struct{ I int }](),
	} {
		ei := Indirect(e)
		t.Logf("type: [%40v] value: [%+40v]", ei.T, ei.V)
	}
}
