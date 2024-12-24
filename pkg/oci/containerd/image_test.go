package containerd

import (
	"testing"

	"github.com/Juminiy/kube/pkg/util"
)

func TestClient_ImageList(t *testing.T) {
	images, err := _cli.ImageList()
	util.Must(err)
	t.Log(GreenPretty(images))
}

func TestClient_ImagePull(t *testing.T) {
	image, err := _cli.ImagePull("192.168.31.242:8662/library/hello-world@sha256:d37ada95d47ad12224c205a938129df7a3e52345828b4fa27b03a98825d1e2e7")
	util.Must(err)
	t.Log(GreenPretty(image))
}
