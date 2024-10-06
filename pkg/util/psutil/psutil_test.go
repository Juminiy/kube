package psutil

import (
	"github.com/Juminiy/kube/pkg/util"
	"testing"
)

func TestMarshal(t *testing.T) {
	t.Log(util.Bytes2StringNoCopy(MarshalIndent()))
}
