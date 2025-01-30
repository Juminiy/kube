package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"testing"
)

// +passed
func TestClient_ListContainers(t *testing.T) {
	containers, err := _cli.ListContainers()
	util.SilentPanic(err)
	t.Log(safe_json.Pretty(containers))
}

// +passed
func TestClient_ListContainerIds(t *testing.T) {
	ids, err := _cli.ListContainerIds()
	util.SilentPanic(err)
	t.Log(ids)
}

// +passed
func TestClient_ListContainerNames(t *testing.T) {
	names, err := _cli.ListContainerNames()
	util.SilentPanic(err)
	t.Log(names)
}
