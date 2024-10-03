package docker_api

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	kubedockertypes "github.com/Juminiy/kube/pkg/image_api/docker_api/types"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"io"
	"os"
	"testing"
)

var (
	imageRef = kubedockertypes.ImageRef{
		Registry:   "192.168.31.242:8662",
		Project:    "library",
		Repository: "hello-world",
		Tag:        "latest",
	}
)

// +passed
func TestClient_ExportImage(t *testing.T) {
	initFunc()
	imageRC, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	imageBytes, err := io.ReadAll(imageRC)
	defer util.SilentCloseIO("image read error", imageRC)

	stdlog.InfoF("size of image amd64 %s is: %s", imageRef.String(), util.BytesOf(imageBytes))

	err = util.TarIOReader2File(imageRC, testTarGzPath)
	util.SilentPanic(err)
	stdlog.InfoF("success save tar file: %s", testTarGzPath)
}

// +failed
func TestClient_ImportImage(t *testing.T) {
	initFunc()
	imageFile, err := os.Open(testTarGzPath)
	util.SilentPanic(err)
	_, err = testNewClient.ImportImage(imageRef.String(), imageFile)
	util.SilentPanic(err)
}

// +passed
func TestClient_ExportImageImportImage(t *testing.T) {
	initFunc()
	imageRC, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	//imageBytes, err := io.ReadAll(imageRC)
	defer util.SilentCloseIO("image read error", imageRC)
	//stdlog.InfoF("size of image amd64 %s is: %d", imageRef.String(), len(imageBytes))

	newImageRef := kubedockertypes.ImageRef{
		Registry:   "192.168.31.242:8662",
		Project:    "library",
		Repository: "hello-world",
		Tag:        "wiwi-x",
	}
	importResp, err := testNewClient.ImportImage(newImageRef.String(), imageRC)
	defer util.SilentCloseIO("import resp", importResp)
	stdlog.Info(docker_internal.GetStatusFromImagePushResp(importResp))
	util.SilentPanic(err)
}
