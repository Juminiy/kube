// Package docker_inst/client.go was generated by codegen, please fix its package dependency, but do not modify its functionality
package docker_inst

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"io"
)

func CreateImageTag(_string string, _string2 string) (io.ReadCloser, error) {
	return _dockerClient.CreateImageTag(_string, _string2)
}

func ExportContainer(_string string) (io.ReadCloser, error) {
	return _dockerClient.ExportContainer(_string)
}

func ExportImage(_string string) (io.ReadCloser, error) {
	return _dockerClient.ExportImage(_string)
}

func GC(varLenfunc ...util.Func) {
	_dockerClient.GC(varLenfunc...)
}

func HostImageStorageGC(varLenhostImageGCFunc ...docker_api.HostImageGCFunc) {
	_dockerClient.HostImageStorageGC(varLenhostImageGCFunc...)
}

func ImportContainer(_string string) (io.ReadCloser, error) {
	return _dockerClient.ImportContainer(_string)
}

func ImportImage(_string string, reader io.Reader) (io.ReadCloser, error) {
	return _dockerClient.ImportImage(_string, reader)
}

func InspectImage(_string string) (types.ImageInspect, error) {
	return _dockerClient.InspectImage(_string)
}

func ListContainerFullIds() ([]string, error) {
	return _dockerClient.ListContainerFullIds()
}

func ListContainerIds() ([]string, error) {
	return _dockerClient.ListContainerIds()
}

func ListContainerNames() ([]string, error) {
	return _dockerClient.ListContainerNames()
}

func ListContainers() ([]types.Container, error) {
	return _dockerClient.ListContainers()
}

func RegistryAuth(authConfig *registry.AuthConfig) (string, error) {
	return _dockerClient.RegistryAuth(authConfig)
}

func WithContext(context context.Context) *docker_api.Client {
	return _dockerClient.WithContext(context)
}

func WithPage(page *util.Page) *docker_api.Client {
	return _dockerClient.WithPage(page)
}

func WithRegistryAuth(authConfig *registry.AuthConfig) *docker_api.Client {
	return _dockerClient.WithRegistryAuth(authConfig)
}
