package docker_inst

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api"
	dockercli "github.com/docker/docker/client"
)

var (
	_dockerClient *dockercli.Client

	_hostURL string
	_version string
	_context context.Context
)

func Init() {
	docker_api.New()
}
