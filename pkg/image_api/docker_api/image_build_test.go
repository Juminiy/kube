package docker_api

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"os"
	"path/filepath"
	"testing"
)

func TestClient_BuildImageV3(t *testing.T) {
	fptr, err := os.Open(filepath.Join(_testTar, "netconn.tar"))
	util.Must(err)
	defer util.SilentCloseIO("tar fileptr", fptr)
	resp, err := _cli.BuildImageV3(util.TODOContext(), fptr, types.ImageBuildOptions{
		Tags:         []string{"library/netconn:v1.0"},
		CPUSetCPUs:   "",
		CPUSetMems:   "",
		CPUShares:    0,
		CPUQuota:     0,
		CPUPeriod:    0,
		Memory:       0,
		MemorySwap:   0,
		CgroupParent: "",
		NetworkMode:  "",
		ShmSize:      0,
		Dockerfile:   "Dockerfile",
		Ulimits:      nil,
		BuildArgs: map[string]*string{
			"OS":      util.NewString("linux"),
			"ARCH":    util.NewString("amd64"),
			"GOPROXY": util.NewString("https://goproxy.cn,direct"),
		},
		Labels: map[string]string{
			"for-test": "true",
			"user-id":  "chisatox",
		},
		Squash:      false,
		CacheFrom:   nil,
		SecurityOpt: nil,
		ExtraHosts:  nil,
		Target:      "",
		Platform:    PlatformLinuxAmd64,
	})
	util.Must(err)
	t.Log(resp)
}
