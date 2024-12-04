package docker_api

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	kubedockerclicommand "github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal/cli/command"
	kubedockertypes "github.com/Juminiy/kube/pkg/image_api/docker_api/types"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/moby/sys/sequential"
	"io"
	"testing"
)

var (
	imageRef = kubedockertypes.ImageRef{
		Registry:   harborAddr,
		Project:    "library",
		Repository: "hello",
		Tag:        "v1.0",
	}
	newImageRef = kubedockertypes.ImageRef{
		Registry:   harborAddr,
		Project:    "library",
		Repository: "hello",
		Tag:        "v2.0",
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

	//err = util.TarIOReaderToFileV2(imageRC, testTarGzPath)
	err = kubedockerclicommand.CopyToFile(testTarGzPath, imageRC)
	util.SilentPanic(err)
	stdlog.InfoF("success save tar file: %s", testTarGzPath)
}

// +failed
func TestClient_ImportImage(t *testing.T) {
	initFunc()
	var input io.Reader
	file, err := sequential.Open(testTarGzPath)
	util.Must(err)
	input = file
	_, err = testNewClient.ImportImage(newImageRef.String(), input)
	util.SilentPanic(err)
	util.SilentCloseIO("file ptr", file)
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
	importResp, err := testNewClient.ImportImage(newImageRef.String(), imageRC)
	defer util.SilentCloseIO("import resp", importResp)
	stdlog.Info(docker_internal.GetStatusFromImagePushResp(importResp))
	util.SilentPanic(err)
}
