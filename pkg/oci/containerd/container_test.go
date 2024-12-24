package containerd

import (
	"testing"

	"github.com/Juminiy/kube/pkg/util"
)

func TestClient_ContainerList(t *testing.T) {
	crts, err := _cli.ContainerList()
	util.Must(err)
	t.Log(GreenPretty(crts))
}
