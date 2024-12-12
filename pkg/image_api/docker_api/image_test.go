package docker_api

import (
	"encoding/base64"
	kubedockerclicommand "github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal/cli/command"
	kubedockertypes "github.com/Juminiy/kube/pkg/image_api/docker_api/types"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/moby/sys/sequential"
	"io"
	"os"
	"testing"
	"time"
)

var (
	imageRefV10 = kubedockertypes.ImageRef{
		Registry:   harborAddr,
		Project:    "library",
		Repository: "hello",
		Tag:        "v1.0",
	}
	imageRegV30 = kubedockertypes.ImageRef{
		Registry:   harborAddr,
		Project:    "library",
		Repository: "hello",
		Tag:        "v3.0",
	}
)

// +passed
func TestClient_ExportImage(t *testing.T) {
	initFunc()
	resp, err := testNewClient.ExportImage(imageRefV10.String())
	util.Must(err)

	if !resp.NotFoundInRegistry {
		t.Log(resp.ImagePulledInfo)
	}

	err = kubedockerclicommand.CopyToFile(testTarGzPath, resp.ImageFileReader)
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
	resp, err := testNewClient.ImportImage(imageRegV30.String(), input)
	util.SilentPanic(err)
	t.Log(resp)
}

func TestClient_pushImageV2(t *testing.T) {
	initFunc()
	resp, err := testNewClient.pushImageV2(imageRegV30.String())
	util.Must(err)
	t.Log(resp)
}

func TestClient_ImportImageV2(t *testing.T) {
	initFunc()
	var input io.Reader
	file, err := sequential.Open(testTarGzPath)
	util.Must(err)
	//defer util.SilentCloseIO("file ptr", file)
	input = file
	resp, err := testNewClient.ImportImageV2(imageRegV30.String(), input)
	util.SilentPanic(err)
	t.Log(resp)
}

// +passed
func TestClient_ExportImageImportImage(t *testing.T) {
	initFunc()
	exportResp, err := testNewClient.ExportImage(imageRefV10.String())
	if err != nil {
		panic(err)
	}

	//imageBytes, err := io.ReadAll(imageRC)
	defer util.SilentCloseIO("image read error", exportResp.ImageFileReader)
	//stdlog.InfoF("size of image amd64 %s is: %d", imageRef.String(), len(imageBytes))
	importResp, err := testNewClient.ImportImage(imageRegV30.String(), exportResp.ImageFileReader)
	util.SilentPanic(err)
	t.Log(importResp)
}

func TestFakeLogin(t *testing.T) {
	//YWRtaW46YnVwdC5oYXJib3JANjY2
	t.Log(base64.StdEncoding.EncodeToString([]byte(harborAuthUsername + ":" + harborAuthPassword)))
}

func TestClient_ImportImageV3(t *testing.T) {
	initFunc()
	var input io.Reader
	file, err := sequential.Open(testTarGzPath)
	util.Must(err)
	//defer util.SilentCloseIO("file ptr", file)
	input = file
	resp, err := testNewClient.ImportImageV3(imageRegV30.String(), input)
	util.SilentPanic(err)
	t.Log(resp)
}

func TestClient_ExportImageV2(t *testing.T) {
	initFunc()
	resp, err := testNewClient.ExportImageV2(imageRegV30.String())
	util.Must(err)

	err = kubedockerclicommand.CopyToFile(testTarGzPath, resp.ImageFileReader)
	util.Must(err)
	t.Log(safe_json.Pretty(resp))
}

func TestClient_BuildImage(t *testing.T) {
	cli := initFunc2()
	cli.WithProject("library")
	fptr, err := os.Open(testTarBuildPath)
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	resp, err := cli.BuildImage(fptr, "jammy-env:v1.9")
	util.Must(err)
	t.Log(resp)
}

func TestClient_BuildImageWithCancel(t *testing.T) {
	cli := initFunc2()
	cli.WithProject("library")
	fptr, err := os.Open(testTarBuildTimeout)
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	ctx := util.TODOContext()
	resp, cancelFunc, err := cli.BuildImageWithCancel(ctx, fptr, "timeout:v1.0")
	util.Must(err)
	t.Log(resp)

	time.Sleep(util.TimeSecond(10))
	if cancelFunc != nil {
		//(*cancelFunc)()
	}
}

func TestClient_BuildImageV2(t *testing.T) {
	cli := initFunc2()
	cli.WithProject("library")
	fptr, err := os.Open(testTarBuildTimeout)
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	resp, err := cli.BuildImageV2(fptr, "timeout:v1.0")
	util.Must(err)
	t.Log(resp)
}
