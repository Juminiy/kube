package containerd

import (
	"testing"

	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
)

var _cli *Client

func init() {
	var _cliErr error
	_cli, _cliErr = New()
	util.Must(_cliErr)
}

func TestNew(t *testing.T) {
	serverUUID, err := _cli.Server()
	util.Must(err)
	t.Log(serverUUID)
}

var GreenPretty = func(i any) string {
	return util.GreenAny(safe_json.Pretty(i))
}

var RedPretty = func(i any) string {
	return util.RedAny(safe_json.Pretty(i))
}
