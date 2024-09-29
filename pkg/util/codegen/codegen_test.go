package codegen

import (
	kubeinternal "github.com/Juminiy/kube/pkg/internal"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util/codegen/example"
	"reflect"
	"testing"
)

/*
func TestManifest_Do_DockerInst(t *testing.T) {
	dockerInst := Manifest{
		DstFilePath:           filepath.Join(workPath, "pkg", "image_api", "docker_api", "docker_inst", "client.go"),
		InstanceOf:            &docker_api.Client{},
		UnExportGlobalVarName: "_dockerClient",
		GenImport:             true,
		GenVar:                false,
	}
	dockerInst.Do()
}

func TestManifest_Do_HarborInst(t *testing.T) {
	harborInst := Manifest{
		DstFilePath:           filepath.Join(workPath, "pkg", "image_api", "harbor_api", "harbor_inst", "client.go"),
		InstanceOf:            &harbor_api.Client{},
		UnExportGlobalVarName: "_harborClient",
		GenImport:             true,
		GenVar:                false,
	}
	harborInst.Do()
}

func TestManifest_Do_MinioInst(t *testing.T) {
	minioInst := Manifest{
		DstFilePath:           filepath.Join(workPath, "pkg", "storage_api", "minio_api", "minio_inst", "client.go"),
		InstanceOf:            &minio_api.Client{},
		UnExportGlobalVarName: "_minioClient",
		GenImport:             true,
		GenVar:                false,
	}
	minioInst.Do()
}
*/

// +passed +windows
func TestManifest_Do_ExampleStruct(t *testing.T) {
	dstFilePath, err := kubeinternal.GetWorkPath("example", "example_inst", "example_codegen.go")
	if err != nil {
		t.Fatal(err)
	}
	exampleManifest := Manifest{
		DstFilePath: dstFilePath,
		InstanceOf:  &example.ExampleStruct{},
		GenImport:   true,
		GenVar:      true,
	}
	exampleManifest.Do()
}

func TestManifest_Do_differ_slice_and_variable_length_param(t *testing.T) {
	func1 := func(v ...int) {}
	func2 := func(v []int) {}

	typeOfF1 := reflect.TypeOf(func1)
	valueOfF1 := reflect.ValueOf(func1)

	typeOfF2 := reflect.TypeOf(func2)
	valueOfF2 := reflect.ValueOf(func2)

	//isVariableLengthParamIn := func(typ reflect.Type) bool {
	//	// 1.last paramIn and
	//	// 2. must be a slice
	//	return true
	//}

	for i := range typeOfF1.NumIn() {
		paramIn := typeOfF1.In(i)
		stdlog.Info(paramIn)
	}

	for i := range typeOfF2.NumIn() {
		paramIn := typeOfF2.In(i)
		stdlog.Info(paramIn)
	}

	stdlog.Info(typeOfF1, valueOfF1, typeOfF2, valueOfF2)

}
