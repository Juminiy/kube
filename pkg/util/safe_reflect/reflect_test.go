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

func TestHasField3(t *testing.T) {
	slI := []int{1, 2, 3}
	slIPtr := []*int{util.New(1), util.New(2)}
	slT0 := []t0{}
	slT0Ptr := []**t0{}
	t.Log(HasField(slI, "F0", ""))
	t.Log(HasField(slIPtr, "F0", ""))
	t.Log(HasField(slT0, "F0", ""))
	t.Log(HasField(slT0Ptr, "F0", ""))
	util.TestLongHorizontalLine(t)

	t.Log(HasField(&slI, "F0", ""))
	t.Log(HasField(&slIPtr, "F0", ""))
	t.Log(HasField(&slT0, "F0", ""))
	t.Log(HasField(&slT0Ptr, "F0", ""))
	util.TestLongHorizontalLine(t)

	arrI := [5]int{1, 2, 3}
	arrPtr := [5]*int{util.New(1), util.New(2)}
	arrT0 := [5]t0{}
	arrT0Ptr := [5]**t0{}
	t.Log(HasField(arrI, "F0", ""))
	t.Log(HasField(arrPtr, "F0", ""))
	t.Log(HasField(arrT0, "F0", ""))
	t.Log(HasField(arrT0Ptr, "F0", ""))
	util.TestLongHorizontalLine(t)

	t.Log(HasField(&arrI, "F0", ""))
	t.Log(HasField(&arrPtr, "F0", ""))
	t.Log(HasField(&arrT0, "F0", ""))
	t.Log(HasField(&arrT0Ptr, "F0", ""))
	util.TestLongHorizontalLine(t)

}

func TestHasField4(t *testing.T) {
	m := map[string]any{"F0": 1}
	t.Log(HasField(m, "F0", 0))
	t.Log(HasField(&m, "F0", ""))
	t.Log(HasField(m, "F1", 0))
	t.Log(HasField(&m, "F1", 0))
	util.TestLongHorizontalLine(t)

	m1 := map[string]int{"F0": 1}
	t.Log(HasField(m1, "F0", 0))
	t.Log(HasField(&m1, "F0", ""))
	t.Log(HasField(m1, "F1", 0))
	t.Log(HasField(&m1, "F1", 0))
	util.TestLongHorizontalLine(t)

	m2 := map[string]*int{"F0": util.New(1)}
	t.Log(HasField(m2, "F0", 0))
	t.Log(HasField(&m2, "F0", ""))
	t.Log(HasField(m2, "F1", 0))
	t.Log(HasField(&m2, "F1", 0))
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

func TestSetField3(t *testing.T) {
	m := map[string]any{"F0": 1}
	SetField(m, "F0", 111)
	t.Log(m)
	SetField(&m, "F0", "xxx")
	t.Log(m)
	SetField(m, "F1", 0)
	t.Log(m)
	SetField(&m, "F1", 0)
	t.Log(m)
	util.TestLongHorizontalLine(t)

	m1 := map[string]int{"F0": 1}
	SetField(m1, "F0", 222)
	t.Log(m1)
	SetField(&m1, "F0", "kkk")
	t.Log(m1)
	SetField(m1, "F1", 0)
	t.Log(m1)
	SetField(&m1, "F1", 0)
	t.Log(m1)
	util.TestLongHorizontalLine(t)

	// need to support
	m2 := map[string]*int{"F0": util.New(1)}
	SetField(m2, "F0", 333)
	t.Log(m2)
	SetField(&m2, "F0", "nnn")
	t.Log(m2)
	SetField(m2, "F1", 0)
	t.Log(m2)
	SetField(&m2, "F1", 0)
	t.Log(m2)
}

func TestHasFields(t *testing.T) {
	t.Log(HasFields(t0{}, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "ccc",                 // has name and type is
		"F1": "1",                   // has name but type not
	}))

	t.Log(HasFields(t0{}, map[string]any{
		"F0": "ccc", // has name and type is
		"F1": 1,     // has name but type is
	}))
}

func TestHasFields2(t *testing.T) {
	t.Log(HasFields([]t0{}, map[string]any{
		"F0": "ccc", // has name and type is
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

	SetFields(&tval, map[string]any{
		"F2": util.NewString("vvv"), // no name
		"F0": "vvv",                 // has name and type is
		"F1": 222,                   // has name and type is
	})
	t.Log(tval)
}

func TestSetFields2(t *testing.T) {
	tval := []t0{{}}

	SetFields(tval, map[string]any{
		"F0": "xxx", // has name and type is
		"F1": 888,   // has name and type is
	})
	t.Log(tval)

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
		t.Logf("%12s value_type(%6s) CanSet(%1s) CanAddr(%1s)", desc, trimTyp(rValueType(v).String()), yn(v.CanSet()), yn(v.CanAddr()))
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

	util.TestLongHorizontalLine(t)

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

func TestVVV(t *testing.T) {
	// package `of api`, include at least once
	//reflect.ArrayOf()
	//reflect.ChanOf()
	//reflect.FuncOf()
	//reflect.MapOf()
	//reflect.SliceOf()
	//reflect.StructOf()
	//reflect.TypeOf()
	//reflect.ValueOf()

	// package `make api`, include at least once
	//reflect.MakeChan()
	//reflect.MakeFunc()
	//reflect.MakeMapWithSize()
	//reflect.MakeMap()
	//reflect.MakeSlice()

	//
	//reflect.New()
	//reflect.Append()
	//reflect.AppendSlice()
	//reflect.Copy()
	//reflect.NewAt()
	//reflect.PointerTo()
	//reflect.Select()
	//reflect.SliceAt()
	//reflect.TypeFor()
	//reflect.Zero()
	//reflect.DeepEqual()
	//reflect.Swapper()
	//reflect.VisibleFields()
	//reflect.Indirect()
}

func TestGetTags(t *testing.T) {

	t.Log("Struct")
	t.Log(GetTags(t0{}, "gorm", "column"))

	t.Log(GetTags(&t0{}, "gorm", "column"))

	t.Log(GetTags(t1{}, "gorm", "column"))

	t.Log(GetTags(&t1{}, "gorm", "column"))
	util.TestLongHorizontalLine(t)

	t.Log("Array")
	t.Log(GetTags([2]t0{}, "gorm", "column"))

	t.Log(GetTags(&[2]t0{}, "gorm", "column"))
	t.Log(GetTags(&[2]*t0{}, "gorm", "column"))
	t.Log(GetTags(&[2]**t0{}, "gorm", "column"))

	t.Log(GetTags([3]t1{}, "gorm", "column"))

	t.Log(GetTags(&[5]t1{}, "gorm", "column"))
	t.Log(GetTags(&[5]*t1{}, "gorm", "column"))
	t.Log(GetTags(&[5]**t1{}, "gorm", "column"))
	util.TestLongHorizontalLine(t)

	t.Log("Slice")
	t.Log(GetTags([]t0{}, "gorm", "column"))

	t.Log(GetTags(&[]t0{}, "gorm", "column"))
	t.Log(GetTags(&[]*t0{}, "gorm", "column"))
	t.Log(GetTags(&[]**t0{}, "gorm", "column"))

	t.Log(GetTags([]t1{}, "gorm", "column"))

	t.Log(GetTags(&[]t1{}, "gorm", "column"))
	t.Log(GetTags(&[]*t1{}, "gorm", "column"))
	t.Log(GetTags(&[]**t1{}, "gorm", "column"))
	util.TestLongHorizontalLine(t)
}
