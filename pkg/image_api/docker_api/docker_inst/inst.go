// global var
package docker_inst

import (
	"context"
	"github.com/Juminiy/kube/pkg/image_api/docker_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	dockercli "github.com/docker/docker/client"
	"net/http"
)

// global config
var (
	_hostURL string
	_version string
)

// global var
var (
	_dockerClient  *docker_api.Client
	_docketContext context.Context
)

func Init() {
	var dockerClientError error
	_dockerClient, dockerClientError = docker_api.New(_hostURL, _version)
	if dockerClientError != nil {
		stdlog.ErrorF("docker client version: %s connect to host: %s error: %s", _version, _hostURL, dockerClientError.Error())
		return
	}
	if _docketContext != nil {
		_dockerClient.WithContext(_docketContext)
	}
}

var (
	_httpClient *http.Client
	_dockerOpts []dockercli.Opt
)

func InitWithOpts() {
	var dockerClientError error
	_dockerClient, dockerClientError = docker_api.NewWithOpts(
		_hostURL, _version,
		_httpClient,
		_dockerOpts...,
	)
	if dockerClientError != nil {
		stdlog.ErrorF("docker client version: %s connect to host: %s error: %s", _version, _hostURL, dockerClientError.Error())
		return
	}
	if _docketContext != nil {
		_dockerClient.WithContext(_docketContext)
	}
}
