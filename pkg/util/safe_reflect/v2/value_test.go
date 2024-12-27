package safe_reflectv2

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func testSetLog(t *testing.T, setFunc func(src, dst any), src, dst any) {
	dstSetBefore := cast.ToString(dst)
	setFunc(src, dst)
	t.Logf("before: %s, after: %s", dstSetBefore, cast.ToString(dst))
}

func TestSet(t *testing.T) {
	var dst int
	testSetLog(t, Set, 111, dst)
	testSetLog(t, Set, 222, &dst)
	testSetLog(t, Set, 333, &dst)
}

func TestSetLike(t *testing.T) {
	var dst int
	testSetLog(t, SetLike, "111", dst)
	testSetLog(t, SetLike, "222", &dst)
	testSetLog(t, SetLike, "333", &dst)
}

func TestIndirect(t *testing.T) {
	var i = 1
	d := Direct(wrapPtr(&i, 10))
	ind := d.indirect()
	ind.SetI(20)
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
