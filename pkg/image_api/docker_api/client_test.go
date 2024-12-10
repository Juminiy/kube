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

func initFunc2() *Client {
	cli, err := New("tcp://192.168.31.242:2375", "1.47")
	util.Must(err)
	cli.WithRegistryAuth(&registry.AuthConfig{
		Username:      "admin",
		Password:      "Harbor12345",
		ServerAddress: "192.168.31.242:8662",
	})
	return cli
}
