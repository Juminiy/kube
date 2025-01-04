package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestTv_Values(t *testing.T) {
	for _, i := range []any{
		&ExportStruct{Int: 666, Float: 888, String: "114514"},
		&[]*ExportStruct{
			{Int: 666, Float: 888, String: "114514"},
		},
		&[1]*ExportStruct{
			{Int: 666, Float: 888, String: "114514"},
		},
		&map[string]any{
			"Languages": []string{"English", "Chinese", "Japanese"},
			"Nation":    "China",
			"Tail":      false,
		},
		&[1]*map[string]any{
			{
				"Languages": []string{"English", "Chinese", "Japanese"},
				"Nation":    "China",
				"Tail":      false,
			},
		},
		&[]*map[string]any{
			{
				"Languages": []string{"English", "Chinese", "Japanese"},
				"Nation":    "China",
				"Tail":      false,
			},
		},
	} {
		t.Log(Indirect(&i).Values())
	}
}

func TestTv_CallMethod(t *testing.T) {
	for _, i := range []any{
		ExportStruct{},
		&ExportStruct{},
		util.New(&ExportStruct{}),
		[]ExportStruct{{}},
		[]*ExportStruct{{}},
		&[]*ExportStruct{{}},
		[1]ExportStruct{{}},
		[1]*ExportStruct{{}},
		&[1]*ExportStruct{{}},
	} {
		d := Direct(i)
		t.Log(d.Type)
		t.Log(d.CallMethod("StructMethod", nil))
		t.Log(d.CallMethod("PointerMethod", nil))
		util.TestLongHorizontalLine(t)
	}
}

func TestV_SetField(t *testing.T) {
	for _, i := range []any{
		ExportStruct{},
		&ExportStruct{},
		util.New(&ExportStruct{}),
		[]ExportStruct{{}},
		[]*ExportStruct{{}},
		&[]*ExportStruct{{}},
		[1]ExportStruct{{}},
		[1]*ExportStruct{{}},
		&[1]*ExportStruct{{}},
		map[string]any{},
		map[string]int{},
		map[string]string{},
	} {
		Indirect(i).SetField(map[string]any{
			"Int":    uint(666),
			"Float":  "114.514",
			"String": 1919810,
		})
		t.Log(Direct(i).Type, Pretty(i))
	}
}

func TestV_SetI(t *testing.T) {
	var slice = []ExportStruct{{}}
	Indirect(slice).SliceSet(0, ExportStruct{
		Int:    666,
		Float:  114.514,
		String: "1919810",
	})
	t.Log(slice)
	var array = [1]ExportStruct{{}}
	Indirect(&array).ArraySet(0, ExportStruct{
		Int:    666,
		Float:  114.514,
		String: "1919810",
	})
	t.Log(array)
}
