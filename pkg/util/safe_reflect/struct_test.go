package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

type t0 struct {
	F0 string `gorm:"column:user_name;type:varchar(128);comment:user's name, account's name" json:"f0,omitempty" app:"name"`
	F1 int    `app:"i"`
}

// +passed
func TestTypVal_ParseStructTag(t *testing.T) {
	// no pointer
	tagMap := Of(t0{}).StructParseTag("gorm")
	t.Log(tagMap)
	t.Log(tagMap.ParseGetVal("F0", "column"))

	// pointer
	pTagMap := Of(&t0{}).StructParseTag("gorm")
	t.Log(pTagMap)
	t.Log(pTagMap.ParseGetVal("F0", "column"))
}

// +passed
func TestTypVal_StructSet(t *testing.T) {
	tvv := t0{}

	src := t0{F0: "no pointer", F1: 69} // no pointer
	srcPtr := &src                      // p
	srcPPtr := &srcPtr                  // pp

	Of(tvv).StructSet(src) // no pointer
	t.Log(tvv)
	Of(tvv).StructSet(srcPtr) // p
	t.Log(tvv)
	Of(tvv).StructSet(srcPPtr) // pp
	t.Log(tvv)
	Of(tvv).StructSet(&srcPPtr) // ppp
	t.Log(tvv)

	Of(&tvv).StructSet(src) // no pointer
	t.Log(tvv)

	src.F0 = "p"
	Of(&tvv).StructSet(srcPtr) // p
	t.Log(tvv)

	src.F0 = "pp"
	Of(&tvv).StructSet(srcPPtr) // pp
	t.Log(tvv)

	src.F0 = "ppp"
	Of(&tvv).StructSet(&srcPPtr) // ppp
	t.Log(tvv)

}

// +passed
func TestTypVal_StructSetFields(t *testing.T) {
	tvv := t0{F0: "no pointer", F1: 666}

	Of(tvv).StructSetFields(map[string]any{
		"F0": "field F0",          // ok
		"F1": "999",               // value_type mismatch
		"F3": util.New2[t0](t0{}), // no field
	})
	t.Log(tvv)

	Of(&tvv).StructSetFields(map[string]any{
		"F0": util.NewString("field F0"), // value_type indirect
		"F1": "999",                      // value_type mismatch
		"F3": util.New2[t0](t0{}),        // no field
	})
	t.Log(tvv)

	tvvPtr := &tvv
	Of(&tvvPtr).StructSetFields(map[string]any{
		"F0": util.NewString("field F0 pointer"), // value_type indirect
		"F1": 999,                                // value_type ok
		"F3": util.New2[t0](t0{}),                // no field
	})
	t.Log(tvv)
}

func TestStructField(t *testing.T) {
	type t1 struct {
		VName  string
		RIndex int
	}
	type t2 struct {
		How0 int
		May1 *int
	}
	type t3 struct {
		t1
		*t2
		VName  string // same name and same type with embedded struct Field
		RIndex string // same name but diff type with embedded struct Field
		How0   int    // same name and same type with embedded pointer Field
		May1   uint   // same name but diff type with embedded pointer Field
	}

	t3val := t3{
		t1:     t1{VName: "t1.VName", RIndex: 111},
		t2:     &t2{How0: 222, May1: util.New(444)},
		VName:  "t3.VName",
		RIndex: "t3.RIndex",
		How0:   -666,
		May1:   999,
	}

	t3t, t3v := directTV(t3val)
	for tIndex := range t3t.NumField() {
		tFi := t3t.Field(tIndex)
		t.Log(tFi.Name, tFi.Type, tFi.Index)
	}

	util.TestLongHorizontalLine(t)
	for vIndex := range t3v.NumField() {
		vFi := t3v.Field(vIndex)
		t.Log(vFi)
	}

	util.TestLongHorizontalLine(t)
	structFi, ok := t3t.FieldByName("t1")
	if ok {
		t.Log(structFi.Name, structFi.Type, structFi.Index)
	}
	structFi, ok = t3t.FieldByName("VName")
	if ok {
		t.Log(structFi.Name, structFi.Type, structFi.Index)
	}
	structFi = t3t.FieldByIndex([]int{0})
	if ok {
		t.Log(structFi.Name, structFi.Type, structFi.Index)
	}
	structFi = t3t.FieldByIndex([]int{0, 1})
	if ok {
		t.Log(structFi.Name, structFi.Type, structFi.Index)
	}

	util.TestLongHorizontalLine(t)
	vFi := t3v.FieldByName("t1")
	t.Log(vFi)
	vFi = t3v.FieldByName("How0")
	t.Log(vFi)
	vFi = t3v.FieldByIndex([]int{0})
	t.Log(vFi)
	vFi = t3v.FieldByIndex([]int{0, 1}) // show that it is a recursive path t3.[0].[1].[.]...
	t.Log(vFi)
}

func TestStructMake(t *testing.T) {
	t.Log(
		StructMake([]FieldDesc{
			{Name: "I32", Value: int32(1), Tag: `json:"i32"`},
			{Name: "I32", Value: int32(1), Tag: `json:"i32"`},
			{Name: "I32", Value: int32(1), Tag: `json:"i32"`},
		}),
	)

	t.Log(
		StructMake([]FieldDesc{
			{Name: "I32", Value: int32(1), Tag: `json:"i32"`},
			{Name: "F64", Value: 1.23, Tag: `json:"i32"`},
			{Name: "String", Value: "vvv", Tag: `json:"i32"`},
		}),
	)
}

func TestTypVal_StructParseTag(t *testing.T) {
	type t5 struct {
		Vx  bool
		V1  int    `app:"unique:1;union_unique:0;field:name;follow::"`
		V19 int    `app:"unique:1;union_unique:0;field:name;follow:"`
		V2  string `app:"unique:1;union_unique:1;field:name_part1;follow:ASCII-Colon"`
	}

	tm := Of(t5{}).StructParseTag("gorm")
	t.Logf("no-app: %v", tm.ParseGetVal("V1", "unique")) // no app

	tagMap := Of(t5{}).StructParseTag("app")
	t.Logf("no-field: (%v)", tagMap.ParseGetVal("Vo", "unique"))                  // no field
	t.Logf("no-tag: (%v)", tagMap.ParseGetVal("Vx", "unique"))                    // no tag
	t.Logf("all-ok: (%v)", tagMap.ParseGetVal("V1", "unique"))                    // app-ok field-ok tag-ok
	t.Logf("all-ok: (%v)", tagMap.ParseGetVal("V1", "union_unique"))              // app-ok field-ok tag-ok
	t.Logf("no-tag-key: (%v)", tagMap.ParseGetVal("V1", "field1"))                // no-tag-key
	t.Logf("no-tag-key: (%v)", tagMap.ParseGetVal("V1", "field2"))                // no-tag-key
	t.Logf("all-ok: (%v)", tagMap.ParseGetVal("V1", "field"))                     // app-ok field-ok tag-ok
	t.Logf("all-ok, with(value=Colon): (%v)", tagMap.ParseGetVal("V1", "follow")) // app-ok field-ok tag-ok, with :
	t.Logf("all-ok, with(value=()): (%v)", tagMap.ParseGetVal("V19", "follow"))   // app-ok field-ok tag-ok, with :
}

func TestTypVal_StructParseTag2(t *testing.T) {
	type t5 struct {
		TenantID       bool
		BusinessName   string `gorm:"column:business_name;type:varchar(123);" app:"unique:1;union_unique:0;field:name;follow::"`
		UnionNamePart1 string `gorm:"column:uname_x;type:varchar(123);" app:"unique:0;union_unique:1;field:fvip;index:0;follow::"`
		UnionNamePart2 string `gorm:"column:uname_y;type:varchar(123);" app:"unique:0;union_unique:1;field:fvip;index:1;follow:-"`
	}

	oft5 := Of(t5{})
	appTagMap := oft5.StructParseTag("app")
	gormTagMap := oft5.StructParseTag("gorm")

	selectFields := make([]string, 0, len(appTagMap))
	//uniqueMap := make(map[string][]string, len(appTagMap))
	for fieldName := range appTagMap {
		uniqueOk := appTagMap.ParseGetVal(fieldName, "unique") == "1"
		unionUniqueOk := appTagMap.ParseGetVal(fieldName, "union_unique") == "1"
		if uniqueOk || unionUniqueOk {
			selectFields = append(selectFields, gormTagMap.ParseGetVal(fieldName, "column"))
		}
		if uniqueOk {
			//uniqueMap[appTagMap.ParseGetVal(fieldName, "field")] =
		}
	}

}

func TestTypVal_StructFieldIndex(t *testing.T) {
	t.Log(Of(t0{}).StructFieldsIndex())
	t.Log(Of(t0{}).StructFieldsType())

	t.Log(Of(&t0{}).StructFieldsIndex())
	t.Log(Of(&t0{}).StructFieldsType())
}

func TestTypVal_StructFieldsValues(t *testing.T) {
	t0Of := Of(t0{
		F0: "avl",
		F1: 111,
	})
	fieldsIndex := t0Of.StructFieldsIndex()

	t.Log(t0Of.StructFieldsValues(fieldsIndex))

	t.Log(t0Of.StructFieldsValues(map[string][]int{
		"Ciallo": nil,
		"Fake":   nil,
		"F0":     nil,
	}))

	t.Log(t0Of.StructFieldValue("F0")) // string

	t.Log(t0Of.StructFieldValue("F1")) // int

	t.Log(t0Of.StructFieldValue("F2")) // nil

	t.Log(t0Of.StructFieldValue("")) // nil
}

func TestTypVal_Struct2Map(t *testing.T) {
	t0Of := Of(t0{
		F0: "avl",
		F1: 111,
	})

	t.Log(t0Of.Struct2Map(map[string]struct{}{}))
	t.Log(t0Of.Struct2Map(map[string]struct{}{
		"F0": {}, "F1": {},
	}))

	t.Log(t0Of.Struct2Map(map[string]struct{}{}))
	t.Log(t0Of.Struct2Map(map[string]struct{}{
		"F0": {}, "F1": {},
	}))
}

func TestTypVal_StructHasFields(t *testing.T) {
	t.Log(Of(t0{}).StructHasFields(map[string]any{
		"F0": 1,
		"F1": "",
	}))

	t.Log(Of(t0{}).StructHasFields(map[string]any{
		"F0": "xsss",
		"F1": 1,
	}))

	t.Log(Of(t0{}).StructHasFields(map[string]any{
		"F2": 1,
		"F3": "",
	}))

	t.Log(Of(t0{}).StructHasFields(map[string]any{
		"F1": 1,
		"F3": "",
	}))

	t.Log(Of(&t0{}).StructHasFields(map[string]any{
		"F0": 1,
		"F1": "",
	}))

	t.Log(Of(&t0{}).StructHasFields(map[string]any{
		"F0": "xsss",
		"F1": 1,
	}))

	t.Log(Of(&t0{}).StructHasFields(map[string]any{
		"F2": 1,
		"F3": "",
	}))

	t.Log(Of(&t0{}).StructHasFields(map[string]any{
		"F1": 1,
		"F3": "",
	}))
}
