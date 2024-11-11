package safe_json

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/valyala/fastjson"
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
