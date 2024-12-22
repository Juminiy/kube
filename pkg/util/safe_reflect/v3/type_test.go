package safe_reflectv3

import "testing"

func TestT_Tag2(t *testing.T) {
	for desc, e := range esvS {
		ei := Indirect(e)
		t.Logf("%40s %10s", desc, ei.Tag2("gorm", "column"))
	}
}
