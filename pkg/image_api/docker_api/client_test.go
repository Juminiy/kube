package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types/registry"
)

const (
	hostURL       = "tcp://192.168.31.242:2375"
	clientVersion = "1.47"
)

var (
	testNewClient, testDockerClientError = New(hostURL, clientVersion)
)

func initFunc() {
	util.SilentPanicError(testDockerClientError)

	testNewClient.WithRegistryAuth(&registry.AuthConfig{
		Username:      "admin",
		Password:      "Harbor12345",
		ServerAddress: "192.168.31.242:8662",
	})
}
