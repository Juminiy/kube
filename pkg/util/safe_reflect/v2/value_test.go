package safe_reflectv2

import (
	"database/sql"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/spf13/cast"
	"testing"
	"time"
)

func testSetLog(t *testing.T, setFunc func(src, dst any), src, dst any) {
	dstSetBefore := cast.ToString(dst)
	setFunc(src, dst)
	dstSetAfter := cast.ToString(dst)
	dstSetTip := "success"
	if dstSetBefore == dstSetAfter {
		dstSetTip = "failed"
	}
	t.Logf("before: %5s, after: %5s, set: %7s", dstSetBefore, dstSetAfter, dstSetTip)
}

func TestSet(t *testing.T) {
	var dst int
	for _, src := range []any{111, 222, "333", util.New(444), util.New("555")} {
		testSetLog(t, Set, src, &dst)
	}
}

func TestSetLike(t *testing.T) {
	var dst int
	for _, src := range []any{111, 222, "333", util.New(444), util.New("555")} {
		testSetLog(t, SetLike, src, &dst)
	}
}

func TestIndirect(t *testing.T) {
	var i = 111
	Indirect(wrapPtr(&i, 10)).SetI(222)
	t.Log(i)
	Indirect(wrapPtr(&i, 10)).SetILike(333)
	t.Log(i)
}

type Intf0 interface {
	Miao()
}

type Intf1 interface {
	Intf0
}

type Neko string

func (n Neko) Miao() {
	fmt.Printf("neko: %s miao\n", n)
}

func TestSetNilIFace(t *testing.T) {
	var intf0 Intf0
	var intf1 Intf1
	testSetLog(t, Set, Neko("Cow"), intf0)
	testSetLog(t, Set, Neko("Cow"), intf1)
	testSetLog(t, SetLike, Neko("Raccoon"), intf0)
	testSetLog(t, SetLike, Neko("Raccoon"), intf1)
}

func TestSetValueIFace(t *testing.T) {
	var intf0 Intf0 = Neko("Benin")
	var intf1 Intf1 = Neko("Calico")
	testSetLog(t, Set, Neko("Cow"), intf0)
	testSetLog(t, Set, Neko("Cow"), intf1)
	testSetLog(t, SetLike, Neko("Raccoon"), intf0)
	testSetLog(t, SetLike, Neko("Raccoon"), intf1)
}

func TestSetNil(t *testing.T) {
	var i32 int32 = 1114
	t.Log(i32)
	SetLike(nil, &i32)
	t.Log(i32)
}

func TestSetSql(t *testing.T) {
	var i32 sql.NullInt32
	t.Log(i32)
	SetLike(1123, &i32)
	t.Log(i32)
	SetLike(nil, &i32)
	t.Log(i32)
}

func TestSetLikeFull(t *testing.T) {
	var t32 sql.NullTime
	t.Log(t32)
	SetLike("2024-09-19 10:22:11", &t32)
	t.Log(t32)
	SetLike(nil, &t32)
	t.Log(t32)
}

func TestSetLikeMultiFull(t *testing.T) {
	var t32 ***sql.NullTime
	t.Log(t32)
	SetLike(time.Now(), &t32)
	t.Log(t32, ***t32)
	SetLike(nil, &t32)
	t.Log(t32, ***t32)
}
