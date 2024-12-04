package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
)

const (
	hostIP              = "10.112.121.243"
	dockerPort          = "2375"
	dockerAddr          = "tcp://" + hostIP + ":" + dockerPort
	dockerClientVersion = "1.43"
	harborPort          = "8111"
	harborAddr          = hostIP + ":" + harborPort
	harborAuthUsername  = "admin"
	harborAuthPassword  = "bupt.harbor@666"
)

var (
	testNewClient, testDockerClientError = New(dockerAddr, dockerClientVersion)
)

func initFunc() {
	util.SilentPanic(testDockerClientError)

	testNewClient.WithRegistryAuth(&registry.AuthConfig{
		Username:      harborAuthUsername,
		Password:      harborAuthPassword,
		ServerAddress: harborAddr,
	})
}
