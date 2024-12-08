package main

import (
	"github.com/Juminiy/kube/pkg/storage_api/minio_api"
	"github.com/Juminiy/kube/pkg/util/codegen"
	"path/filepath"
)

func main() {
	minioCodegen()
}

func minioCodegen() {
	minioInst := codegen.Manifest{
		DstFilePath:           filepath.Join("minio_inst", "client.go"),
		InstanceOf:            &minio_api.Client{},
		UnExportGlobalVarName: "_minioClient",
		GenImport:             true,
		GenVar:                false,
	}
	minioInst.Do()
}
