package docker_api

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
)

const (
	hostURL       = "tcp://docker.local:2375"
	clientVersion = "1.43"
)

var (
	testNewClient, testDockerClientError = New(hostURL, clientVersion)
)

func TestClient_ListContainers(t *testing.T) {
	util.SilentPanicError(testDockerClientError)
	containers, err := testNewClient.ListContainers()
	util.SilentPanicError(err)
	str := strings.Builder{}
	encoder := json.NewEncoder(&str)
	encoder.SetIndent(util.JSONMarshalPrefix, util.JSONMarshalIndent)
	err = encoder.Encode(containers)
	util.SilentPanicError(err)
	stdlog.Debug(str.String())
}

func TestClient_ListContainerIds(t *testing.T) {
	util.SilentPanicError(testDockerClientError)
	ids, err := testNewClient.ListContainerIds()
	util.SilentPanicError(err)
	stdlog.Info(ids)

	names, err := testNewClient.ListContainerNames()
	util.SilentPanicError(err)
	stdlog.Info(names)
}
