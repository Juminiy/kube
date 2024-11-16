package safe_json

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/valyala/fastjson"
	"os"
	"testing"
)

func TestFastJSON(t *testing.T) {
	fval, err := fastjson.Parse(`{"k1": "v1", "k2": "v2", "k3": ["v30","v31",32,true,3.33]}`)
	util.Must(err)
	fobj, err := fval.Object()
	util.Must(err)
	t.Log(fobj)
	k2k2 := fval.Get("k2", "k2")
	t.Log(k2k2)
}

func TestExpandFromBytes(t *testing.T) {
	embedFiles, err := internal_api.GetDirFileNames("testdata\\embed")
	util.Must(err)
	for i := range embedFiles {
		testExpand("testdata\\embed\\", "testdata\\expand\\", embedFiles[i])
	}
}

func testExpand(srcDir, dstDir, fileName string) {
	bs, err := os.ReadFile(srcDir + fileName)
	util.Must(err)
	fptr, err := internal_api.OverwriteCreateFile(dstDir + fileName)
	defer util.SilentCloseIO("file ptr", fptr)
	util.Must(err)
	_, err = fptr.Write(ExpandFromBytes(bs, WithIgnoreNull()).Marshal())
	util.Must(err)
}
