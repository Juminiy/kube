package compress

import (
	"github.com/Juminiy/kube/pkg/util"
	"os"
	"testing"
)

func TestBzip2Encode(t *testing.T) {
	bsEnc, err := Bzip2Encode([]byte(util.MagicStr))
	if err != nil {
		t.Error(err)
	}
	t.Logf("\n%s", util.MagicStr)
	t.Logf("\n%s", bsEnc)
}

func TestBzip2Decode(t *testing.T) {
	cialloFile, err := os.Open("testdata/ciallo.txt")
	if err != nil {
		t.Error(err)
	}
	bsDec, err := Bzip2Decode(cialloFile)
	if err != nil {
		t.Error(err)
	}
	util.SilentCloseIO("file ptr", cialloFile)
	t.Logf("%s", bsDec)
}
