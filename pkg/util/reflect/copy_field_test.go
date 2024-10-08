package reflect

import (
	"github.com/Juminiy/kube/pkg/util"
	"strconv"
	"testing"
)

type (
	struct1 struct {
		V1 int
		V2 *struct1
		v3 string
	}
	struct2 struct {
		V1 int
		V2 string
		V3 map[string]struct{}
		v3 string
	}
)

func (s *struct1) String() string {
	return util.StringJoin(", ",
		strconv.Itoa(s.V1),
		util.Ptr2a(s.V2),
		s.v3,
	)
}

func (s *struct2) String() string {
	return util.StringJoin(", ",
		strconv.Itoa(s.V1),
		s.V2,
		"s.V3(map)",
		s.v3,
	)
}

// +passed
func TestCopyFieldValue(t *testing.T) {

	src1 := struct1{V1: 10, V2: &struct1{V1: 20, V2: nil}, v3: "mama"}
	dst1 := &struct1{}
	dst2 := &struct2{V1: 3, V2: "coco", V3: nil}

	src1S, dst1S, dst2S := src1.String(), dst1.String(), dst2.String()

	src1Ptr := &src1
	dst2Ptr := &dst2
	dst2PtrPtr := &dst2Ptr
	CopyFieldValue(&src1Ptr, &dst1)
	CopyFieldValue(src1, &dst2PtrPtr)

	t.Logf("src1: %s -> %s", src1S, src1.String())
	t.Logf("dst1: %s -> %s", dst1S, dst1.String())
	t.Logf("dst2: %s -> %s", dst2S, dst2.String())
}
