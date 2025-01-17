package safe_reflectv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
	"unsafe"
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

func TestToAnySlice(t *testing.T) {
	for _, v := range []any{
		nil,
		true,
		int(-5), int8(11), int16(32), int32(77), int64(10),
		uint(1), uint8(8), uint16(9), uint32(27), uint64(222),
		uintptr(1122), float32(0.1), float64(1.1),
		complex(10, 20), complex(float64(1), float64(2)),
		[3]uint{10, 11, 12}, [2]int{1, 2}, [3]any{nil, nil, nil},
		make(chan int, 1),
		func() {},
		any(10),
		map[string]any{"k": "v"},
		util.New(10),
		[]uint{1, 2, 3}, []string{"rr", "bb", "aa"}, []any{nil, nil, nil},
		"rrr",
		struct{ V int }{V: 100443},
		unsafe.Pointer(util.New(10)),
	} {
		t.Log(ToAnySlice(v))
	}
}
