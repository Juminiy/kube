package docker_api

import (
	kubedockerclicommand "github.com/Juminiy/kube/pkg/image_api/docker_api/docker_internal/cli/command"
	kubedockertypes "github.com/Juminiy/kube/pkg/image_api/docker_api/types"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"github.com/moby/sys/sequential"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var (
	imageRefV10 = kubedockertypes.ImageRef{
		Registry:   _cfg.Registry.Addr,
		Project:    "library",
		Repository: "hello",
		Tag:        "v1.0",
	}
	imageRegV30 = kubedockertypes.ImageRef{
		Registry:   _cfg.Registry.Addr,
		Project:    "library",
		Repository: "hello",
		Tag:        "v3.0",
	}
	_helloTar = filepath.Join(_testTar, "hello.tar")
)

// +passed
func TestClient_ExportImage(t *testing.T) {
	resp, err := _cli.ExportImage(imageRefV10.String())
	util.Must(err)

	if !resp.NotFoundInRegistry {
		t.Log(resp.ImagePulledInfo)
	}

	err = kubedockerclicommand.CopyToFile(_helloTar, resp.ImageFileReader)
	util.Must(err)
}

// +passed
func TestClient_ImportImage(t *testing.T) {
	var input io.Reader
	file, err := sequential.Open(_helloTar)
	util.Must(err)
	//defer util.SilentCloseIO("file ptr", file)
	input = file
	resp, err := _cli.ImportImage(imageRegV30.String(), input)
	util.SilentPanic(err)
	t.Log(resp)
}

func TestClient_pushImageV2(t *testing.T) {
	resp, err := _cli.pushImageV2(imageRegV30.String())
	util.Must(err)
	t.Log(resp)
}

func TestClient_ImportImageV2(t *testing.T) {
	var input io.Reader
	file, err := sequential.Open(_helloTar)
	util.Must(err)
	//defer util.SilentCloseIO("file ptr", file)
	input = file
	resp, err := _cli.ImportImageV2(imageRegV30.String(), input)
	util.SilentPanic(err)
	t.Log(resp)
}

// +passed
func TestClient_ExportImageImportImage(t *testing.T) {
	exportResp, err := _cli.ExportImage(imageRefV10.String())
	util.Must(err)

	//imageBytes, err := io.ReadAll(imageRC)
	defer util.SilentCloseIO("image read error", exportResp.ImageFileReader)
	//stdlog.InfoF("size of image amd64 %s is: %d", imageRef.String(), len(imageBytes))
	importResp, err := _cli.ImportImage(imageRegV30.String(), exportResp.ImageFileReader)
	util.SilentPanic(err)
	t.Log(importResp)
}

func TestClient_ImportImageV3(t *testing.T) {
	var input io.Reader
	file, err := sequential.Open(_helloTar)
	util.Must(err)
	//defer util.SilentCloseIO("file ptr", file)
	input = file
	resp, err := _cli.ImportImageV3(imageRegV30.String(), input)
	util.SilentPanic(err)
	t.Log(resp)
}

func TestClient_ExportImageV2(t *testing.T) {
	resp, err := _cli.ExportImageV2(imageRegV30.String())
	util.Must(err)

	err = kubedockerclicommand.CopyToFile(_helloTar, resp.ImageFileReader)
	util.Must(err)
	t.Log(safe_json.Pretty(resp))
}

func TestClient_BuildImage(t *testing.T) {
	_cli.WithProject("library")
	fptr, err := os.Open(_helloTar)
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	resp, err := _cli.BuildImage(fptr, "jammy-env:v1.9")
	util.Must(err)
	t.Log(resp)
}

func TestClient_BuildImageWithCancel(t *testing.T) {
	_cli.WithProject("library")
	fptr, err := os.Open(_helloTar)
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	ctx := util.TODOContext()
	resp, cancelFunc, err := _cli.BuildImageWithCancel(ctx, fptr, "timeout:v1.0")
	util.Must(err)
	t.Log(resp)

	time.Sleep(util.TimeSecond(10))
	if cancelFunc != nil {
		//(*cancelFunc)()
	}
}

func TestClient_BuildImageV2(t *testing.T) {
	_cli.WithProject("library")
	fptr, err := os.Open(filepath.Join(_testTar, "netconn.tar"))
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	opt := _cli.BuildImageFavOption("netconn:v1.1")
	opt.BuildArgs = _GoBuildArgs
	//opt.BuildArgs["HTTP_PROXY"] = util.New("192.168.3.37:7890")
	//opt.BuildArgs["HTTPS_PROXY"] = util.New("192.168.3.37:7890")
	resp, err := _cli.BuildImageV4(util.TODOContext(), fptr, opt)
	t.Log(safe_json.Pretty(resp))
	t.Log(err)
}
