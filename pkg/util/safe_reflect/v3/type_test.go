package safe_reflectv3

import "testing"

func TestT_Tag2(t *testing.T) {
	for desc, e := range esvS {
		ei := Indirect(e)
		t.Logf("%40s %10s", desc, ei.Tag2("gorm", "column"))
	}
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
