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
	tagMap := Of(t0{}).ParseStructTag("gorm")
	t.Log(tagMap)
	t.Log(tagMap.ParseGetTagValV("F0", "column"))

	// pointer
	pTagMap := Of(&t0{}).ParseStructTag("gorm")
	t.Log(pTagMap)
	t.Log(pTagMap.ParseGetTagValV("F0", "column"))
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
