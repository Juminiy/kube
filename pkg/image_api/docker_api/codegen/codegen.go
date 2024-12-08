package main

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api"
	"github.com/Juminiy/kube/pkg/util/codegen"
	"path/filepath"
)

func main() {
	dockerCodegen()
}

func dockerCodegen() {
	dockerInst := codegen.Manifest{
		DstFilePath:           filepath.Join("docker_inst", "client.go"),
		InstanceOf:            &docker_api.Client{},
		UnExportGlobalVarName: "_dockerClient",
		GenImport:             true,
		GenVar:                false,
	}
	dockerInst.Do()
}
