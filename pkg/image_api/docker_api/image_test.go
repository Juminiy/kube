package docker_api

import (
	"encoding/base64"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal"
	kubedockerclicommand "github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal/cli/command"
	kubedockertypes "github.com/Juminiy/kube/pkg/image_api/docker_api/types"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/moby/sys/sequential"
	"io"
	"os"
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
		Tag:        "v3.0",
	}
)

// +passed
func TestClient_ExportImage(t *testing.T) {
	initFunc()
	imagePullResp, imageSaveResp, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(os.Stderr, imagePullResp)
	util.Must(err)

	err = kubedockerclicommand.CopyToFile(testTarGzPath, imageSaveResp)
	util.Must(err)
}

// +passed
func TestClient_ImportImage(t *testing.T) {
	initFunc()
	var input io.Reader
	file, err := sequential.Open(testTarGzPath)
	util.Must(err)
	//defer util.SilentCloseIO("file ptr", file)
	input = file
	rc, err := testNewClient.ImportImage(newImageRef.String(), input)
	util.SilentPanic(err)
	t.Log(util.IOGetStr(rc))
}

// +passed
func TestClient_ExportImageImportImage(t *testing.T) {
	initFunc()
	imageTarResp, imageSaveResp, err := testNewClient.ExportImage(imageRef.String())
	if err != nil {
		panic(err)
	}

	//imageBytes, err := io.ReadAll(imageRC)
	defer util.SilentCloseIO("image read error", imageTarResp)
	//stdlog.InfoF("size of image amd64 %s is: %d", imageRef.String(), len(imageBytes))
	importResp, err := testNewClient.ImportImage(newImageRef.String(), imageTarResp)
	util.SilentPanic(err)
	defer util.SilentCloseIO("import resp", importResp)
	stdlog.Info(docker_internal.GetStatusFromImagePushResp(importResp))

	stdlog.InfoF("image save resp: %s", util.IOGetStr(imageSaveResp))
}

func TestFakeLogin(t *testing.T) {
	//YWRtaW46YnVwdC5oYXJib3JANjY2
	t.Log(base64.StdEncoding.EncodeToString([]byte(harborAuthUsername + ":" + harborAuthPassword)))
}
