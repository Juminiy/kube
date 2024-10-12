package safe_reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"reflect"
	"strings"
	"testing"
)

func TestHasField(t *testing.T) {
	t.Log(HasField(t0{}, "F0", "")) // has name and type is
	t.Log(HasField(t0{}, "F1", "")) // has name but type not
	t.Log(HasField(t0{}, "F2", "")) // no name
	t.Log(HasField(t0{}, "F3", 1))  // others
}

func TestSetField(t *testing.T) {
	tval := t0{}
	t.Log(tval)

	SetField(tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(tval, "F3", "set F3") // others
	t.Log(tval)

	SetField(&tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(&tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(&tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(&tval, "F3", "set F3") // others
	t.Log(tval)
}

func TestHasFields(t *testing.T) {
	t.Log(HasFields(t0{}, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": 1,                     // has name but type not
	}))
}

func TestSetFields(t *testing.T) {
	tval := t0{}

	SetFields(tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)
}

func TestHasField2(t *testing.T) {
	sl := []t0{{}, {}, {}}
	t.Log(HasField(sl, "F0", "")) // has name and type is
	t.Log(HasField(sl, "F1", "")) // has name but type not
	t.Log(HasField(sl, "F2", "")) // no name
	t.Log(HasField(sl, "F3", 1))  // others

	// nil
	sl = nil
	t.Log(HasField(sl, "F0", "")) // has name and type is
	t.Log(HasField(sl, "F1", "")) // has name but type not
	t.Log(HasField(sl, "F2", "")) // no name
	t.Log(HasField(sl, "F3", 1))  // others
}

func TestSetField2(t *testing.T) {
	tval := []t0{{}, {}, {}}
	t.Log(tval)

	SetField(tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(tval, "F3", "set F3") // others
	t.Log(tval)

	tval = nil
	SetField(&tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(&tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(&tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(&tval, "F3", "set F3") // others
	t.Log(tval)

	tval = []t0{{}, {}, {}}
	SetField(&tval, "F0", "set F0") // has name and type is
	t.Log(tval)
	SetField(&tval, "F1", "set F1") // has name but type not
	t.Log(tval)
	SetField(&tval, "F2", "set F2") // no name
	t.Log(tval)
	SetField(&tval, "F3", "set F3") // others
	t.Log(tval)
}

func TestHasFields2(t *testing.T) {
	t.Log(HasFields([]t0{}, map[string]any{
		"F0": "ccc", // has name and type is
	}))
}

func TestSetFields2(t *testing.T) {
	tval := []t0{{}}

	SetFields(tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	tval = []t0{{}}
	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	tval = nil
	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)

	tval = []t0{}
	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": util.NewString("mmm"), // has name but type not
	})
	t.Log(tval)
}

var trimTyp = func(s string) string {
	pkgTyp := strings.Split(s, "/")         // aaa/bbb/ccc/ddd.eee
	pkgtyp := strings.Split(pkgTyp[0], ".") // ddd.eee
	return pkgtyp[len(pkgtyp)-1]            // eee
}
var yn = util.YN

func TestHowAssignable(t *testing.T) {
	logFCan := func(v any) {
		dt, dv := directTV(v)
		t.Logf("direct type(%10s) CanSet(%1s) CanAddr(%1s)", trimTyp(dt.String()), yn(dv.CanSet()), yn(dv.CanAddr()))
		it, iv := indirectTV(v)
		t.Logf("underl type(%10s) CanSet(%1s) CanAddr(%1s)", trimTyp(it.String()), yn(iv.CanSet()), yn(iv.CanAddr()))
	}

	var (
		structT0 = t0{}
		arrayI   = [3]int{1, 2, 3}
		sliceI   = []int{4, 5, 6}
		mapST0   = map[string]t0{"k1": {}, "k2": {}}
	)

	// value_type
	// bool
	logFCan(true)
	// int
	logFCan(-1)
	// uint
	logFCan(uint(1))
	// float
	logFCan(22.33)
	// complex
	logFCan(complex(1.14, 5.14))
	// array
	logFCan(arrayI)
	// map
	logFCan(mapST0)
	// slice
	logFCan(sliceI)
	// string
	logFCan("Ciallo~")
	// struct
	logFCan(structT0)

	// pointer_type
	// bool
	logFCan(util.New(true))
	// int
	logFCan(util.New(-1))
	// uint
	logFCan(util.New(uint(1)))
	// float
	logFCan(util.New(22.33))
	// complex
	logFCan(util.New(complex(1.14, 5.14)))
	// array
	logFCan(util.New(arrayI))
	// map
	logFCan(util.New(mapST0))
	// slice
	logFCan(util.New(sliceI))
	// string
	logFCan(util.New("Ciallo~"))
	// struct
	logFCan(util.New(structT0))

	// pointer_pointer_type
	// bool
	logFCan(util.New2(true))
	// int
	logFCan(util.New2(-1))
	// uint
	logFCan(util.New2(uint(1)))
	// float
	logFCan(util.New2(22.33))
	// complex
	logFCan(util.New2(complex(1.14, 5.14)))
	// array
	logFCan(util.New2(arrayI))
	// map
	logFCan(util.New2(mapST0))
	// slice
	logFCan(util.New2(sliceI))
	// string
	logFCan(util.New2("Ciallo~"))
	// struct
	logFCan(util.New2(structT0))
}

func TestHowAssignable2(t *testing.T) {

	logFCan := func(desc string, v reflect.Value) {
		t.Logf("%12s value_type(%6s) CanSet(%1s) CanAddr(%1s)", desc, trimTyp(v.Type().String()), yn(v.CanSet()), yn(v.CanAddr()))
	}
	// ENUM OF
	// (struct, array, slice, map)
	// (bool, int, uint, float, complex, string)
	// (chan, func, interface, pointer, UnsafePointer)

	var (
		structT0 = t0{}
		arrayI   = [3]int{1, 2, 3}
		sliceI   = []int{4, 5, 6}
		mapST0   = map[string]t0{"k1": {}, "k2": {}}
	)

	// no pointer
	// struct-field
	logFCan("struct field", Of(structT0).Val.FieldByName("F0"))

	// array-elem
	logFCan("array elem", Of(arrayI).Val.Index(0))

	// slice-elem
	logFCan("slice elem", Of(sliceI).Val.Index(0))

	// map-key, map-elem
	logFCan("map key", Of(mapST0).Val.MapKeys()[0])
	logFCan("map elem", Of(mapST0).Val.MapIndex(directV("k1")))

	t.Log("----------------------------------------------------------------")

	// pointer
	// struct-field
	logFCan("struct field", IndirectOf(&structT0).Val.FieldByName("F0"))

	// array-elem
	logFCan("array elem", IndirectOf(&arrayI).Val.Index(0))

	// slice-elem
	logFCan("slice elem", IndirectOf(&sliceI).Val.Index(0))

	// map-key, map-elem
	logFCan("map key", IndirectOf(&mapST0).Val.MapKeys()[0])
	logFCan("map elem", IndirectOf(&mapST0).Val.MapIndex(directV("k1")))
}
