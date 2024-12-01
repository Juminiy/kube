package safe_reflectv2

import (
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
