package driver_docker_container

import (
	"bytes"
	dockerapiprovider "github.com/Juminiy/kube/pkg/image_api/docker_api/api_provider"
	"github.com/Juminiy/kube/pkg/util"
)

type BuildImage interface {
	RunBuild() error
}

func NewImageBuilder(api dockerapiprovider.APIProvider, idOrName string) BuildImage {
	return &containerExec{
		APIProvider: api,
		ctx:         util.TODOContext(),
		idOrName:    idOrName,
		stdout:      bytes.NewBuffer(make([]byte, 0, 4*util.Ki)),
		stderr:      bytes.NewBuffer(make([]byte, 0, 4*util.Ki)),
	}
}
