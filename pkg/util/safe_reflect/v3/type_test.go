package safe_reflectv3

import "testing"

func Test_Tag1(t *testing.T) {
	for _, e := range esvS {
		ei := Indirect(e.Value)
		t.Logf("%40s %v", e.Desc, ei.Tag1("gorm"))
	}
}

func TestT_Tag2(t *testing.T) {
	for _, e := range esvS {
		ei := Indirect(e.Value)
		t.Logf("%40s %10s", e.Desc, ei.Tag2("gorm", "column"))
	}
}

func TestT_TagNil(t *testing.T) {
	t.Log(Indirect(nil).Tag1("gorm"))
}

func TestT_NewElem(t *testing.T) {
	for i, val := range []any{
		"str",             // string
		1,                 // int
		uint(2),           // uint
		struct{ I int }{}, // struct{I int}
		[3]int{1, 1, 1},   // int
		[]int{1},          // int
		map[int]string{},  // string
	} {
		tv := Indirect(val).NewElem()
		tv.SetI(struct{ I int }{I: 888})
		tv.SetI((i + 1) * 111)
		t.Logf("%+v", tv.Indirect().I())
	}
}

type named interface {
	Name() string
}

type namedValue struct{}

func (namedValue) Name() string { return `value` }

func TestT_IsEFace(t *testing.T) {
	var strValue any
	strValue = "some value"
	t.Log(Direct(1).IsEFace())
	t.Log(Direct(2.2).IsEFace())
	t.Log(Direct("iammagiboy").IsEFace())
	t.Log(Direct(strValue).IsEFace())
	t.Log(Direct(new(any)).IsEFace())
	t.Log(Direct(named(namedValue{})).IsEFace())
}

func TestIsMapStringAny(t *testing.T) {
	type mapStringAny map[string]any
	type mapStringAny2 = map[string]any
	t.Log(IsMapStringAny(map[string]any{}))
	t.Log(IsMapStringAny(mapStringAny{}))
	t.Log(IsMapStringAny(mapStringAny2{}))

	t.Log(IsMapStringAny(map[string]string{}))
	t.Log(IsMapStringAny(map[string]int{}))
	t.Log(IsMapStringAny(map[string]named{}))

	t.Log(IsMapStringAny(&map[string]any{}))
	t.Log(IsMapStringAny(&mapStringAny{}))
	t.Log(IsMapStringAny(&mapStringAny2{}))

	t.Log(IsMapStringAny(&map[string]string{}))
	t.Log(IsMapStringAny(&map[string]int{}))
	t.Log(IsMapStringAny(&map[string]named{}))
}
