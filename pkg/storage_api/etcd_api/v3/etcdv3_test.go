package etcdv3

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

var _cli *Client
var Enc = safe_json.Pretty
var Dec = safe_json.From

func init() {
	cli, err := New("127.0.0.1:2379")
	util.Must(err)
	_cli = cli
}

func TestClient_Put(t *testing.T) {
	resp, err := _cli.Put("hashlog", "v1")
	util.Must(err)
	t.Log(Enc(resp))
}

func TestClient_Get(t *testing.T) {
	resp, err := _cli.Get("hashlog")
	util.Must(err)
	t.Log(Enc(resp))
}
