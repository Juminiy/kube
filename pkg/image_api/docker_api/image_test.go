package docker_api

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	kubedockertypes "github.com/Juminiy/kube/pkg/image_api/docker_api/types"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"io"
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
	imageRC, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	imageBytes, err := io.ReadAll(imageRC)
	defer util.HandleCloseError("image read error", imageRC)

	stdlog.InfoF("size of image amd64 %s is: %s", imageRef.String(), util.BytesOf(imageBytes))

	err = util.GzipIOReader2File(imageRC, testTarGzPath)
	util.SilentPanicError(err)
	stdlog.InfoF("success save tar file: %s", testTarGzPath)
}

// +failed
func TestClient_ImportImage(t *testing.T) {
	imageFile, err := util.OSOpenFileWithCreate(testTarGZPathExportedByLinux)
	util.SilentPanicError(err)
	_, err = testNewClient.ImportImage(imageRef.String(), imageFile)
	util.SilentPanicError(err)
}

// +passed
func TestClient_ExportImageImportImage(t *testing.T) {
	imageRC, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	//imageBytes, err := io.ReadAll(imageRC)
	defer util.HandleCloseError("image read error", imageRC)
	//stdlog.InfoF("size of image amd64 %s is: %d", imageRef.String(), len(imageBytes))

	newImageRef := kubedockertypes.ImageRef{
		Registry:   "192.168.31.242:8662",
		Project:    "library",
		Repository: "hello-world",
		Tag:        "wiwi-x",
	}
	importResp, err := testNewClient.ImportImage(newImageRef.String(), imageRC)
	defer util.HandleCloseError("import resp", importResp)
	stdlog.Info(docker_internal.GetStatusFromImagePushResp(importResp))
	util.SilentPanicError(err)
}
