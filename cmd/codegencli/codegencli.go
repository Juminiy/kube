package main

import (
	"flag"
	"github.com/Juminiy/kube/pkg/image_api/docker_api"
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api"
	"github.com/Juminiy/kube/pkg/util/codegen"
	"os"
	"path/filepath"
)

func main() {
	flag.StringVar(&genModule, "gen", "", "which module to gen: docker | harbor | minio | all")
	flag.Parse()

	var err error
	workPath, err = os.Getwd()
	if err != nil {
		stdlog.FatalF("os get work directory error: %s", err.Error())
	}
	stdlog.InfoF("current work directory is: %s", workPath)

	switch genModule {
	case "docker":
		dockerCodegen()
	case "harbor":
		harborCodegen()
	case "minio":
		minioCodegen()
	case "all":
		allCodegen()
	default:
		stdlog.Warn(genModule, "do nothing, exit 0")
	}

}

func allCodegen() {
	dockerCodegen()
	harborCodegen()
	minioCodegen()
}

func dockerCodegen() {
	dockerInst := codegen.Manifest{
		DstFilePath:           filepath.Join(workPath, "pkg", "image_api", "docker_api", "docker_inst", "client.go"),
		InstanceOf:            &docker_api.Client{},
		UnExportGlobalVarName: "_dockerClient",
		GenImport:             true,
		GenVar:                false,
	}
	dockerInst.Do()
	tipsPathToFixPackageImport(dockerInst.DstFilePath)
}

func harborCodegen() {
	harborInst := codegen.Manifest{
		DstFilePath:           filepath.Join(workPath, "pkg", "image_api", "harbor_api", "harbor_inst", "client.go"),
		InstanceOf:            &harbor_api.Client{},
		UnExportGlobalVarName: "_harborClient",
		GenImport:             true,
		GenVar:                false,
	}
	harborInst.Do()
	tipsPathToFixPackageImport(harborInst.DstFilePath)
}

func minioCodegen() {
	minioInst := codegen.Manifest{
		DstFilePath:           filepath.Join(workPath, "pkg", "storage_api", "minio_api", "minio_inst", "client.go"),
		InstanceOf:            &minio_api.Client{},
		UnExportGlobalVarName: "_minioClient",
		GenImport:             true,
		GenVar:                false,
	}
	minioInst.Do()
	tipsPathToFixPackageImport(minioInst.DstFilePath)
}

func tipsPathToFixPackageImport(pkgPath string) {
	stdlog.WarnF("please go to dir: %s to fix package import", pkgPath)
}

var (
	workPath string

	genModule string
)
