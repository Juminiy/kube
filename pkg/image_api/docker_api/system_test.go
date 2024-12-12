package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

func TestClient_SystemInfo(t *testing.T) {
	info, err := _cli.SystemInfo()
	util.Must(err)
	t.Log(safe_json.Pretty(info))
}

func TestClient_SystemVersion(t *testing.T) {
	ver, err := _cli.SystemVersion()
	util.Must(err)
	t.Log(safe_json.Pretty(ver))
}

func TestClient_SystemPing(t *testing.T) {
	ping, err := _cli.SystemPing()
	util.Must(err)
	t.Log(safe_json.Pretty(ping))
}
