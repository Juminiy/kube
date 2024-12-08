package main

import (
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/util/codegen"
	"path/filepath"
)

func main() {
	harborCodegen()
}
func harborCodegen() {
	harborInst := codegen.Manifest{
		DstFilePath:           filepath.Join("harbor_inst", "client.go"),
		InstanceOf:            &harbor_api.Client{},
		UnExportGlobalVarName: "_harborClient",
		GenImport:             true,
		GenVar:                false,
	}
	harborInst.Do()
}
