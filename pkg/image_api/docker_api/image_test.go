package docker_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"io"
	"testing"
)

var (
	imageRef = ImageRef{
		Registry:   "192.168.31.242:8662",
		Project:    "library",
		Repository: "nginx",
		Tag:        "stable-alpine3.17",
		absRefStr:  "",
	}
)

func TestClient_ExportImage(t *testing.T) {
	imageRC, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	imageBytes, err := io.ReadAll(imageRC)
	defer util.HandleCloseError("image read error", imageRC)
	stdlog.InfoF("size of image amd64 %s is: %d", imageRef.String(), len(imageBytes))
}

func TestClient_ImportImage(t *testing.T) {
	imageFile, err := util.OSOpenFileWithCreate(testTarGzPath)
	util.SilentPanicError(err)
	_, err = testNewClient.ImportImage(imageRef.String(), imageFile)
	util.SilentPanicError(err)
}
