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
