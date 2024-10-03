package docker_api

import (
	"encoding/json"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"strings"
	"testing"
)

// +passed
func TestClient_ListContainers(t *testing.T) {
	initFunc()
	containers, err := testNewClient.ListContainers()
	util.SilentPanic(err)
	str := strings.Builder{}
	encoder := json.NewEncoder(&str)
	encoder.SetIndent(util.JSONMarshalPrefix, util.JSONMarshalIndent)
	err = encoder.Encode(containers)
	util.SilentPanic(err)
	stdlog.Debug(str.String())
}

// +passed
func TestClient_ListContainerIds(t *testing.T) {
	initFunc()
	ids, err := testNewClient.ListContainerIds()
	util.SilentPanic(err)
	stdlog.Info(ids)
}

// +passed
func TestClient_ListContainerNames(t *testing.T) {
	initFunc()
	names, err := testNewClient.ListContainerNames()
	util.SilentPanic(err)
	stdlog.Info(names)
}
