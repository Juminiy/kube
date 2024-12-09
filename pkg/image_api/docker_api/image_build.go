package docker_api

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"io"
)

func (c *Client) BuildImage(input io.Reader) (types.ImageBuildResponse, error) {
	return c.cli.ImageBuild(c.ctx, input, c.BuildImageFavOption())
}

func (c *Client) BuildImageFavOption() types.ImageBuildOptions {
	return types.ImageBuildOptions{
		Tags:           nil,
		SuppressOutput: false,
		RemoteContext:  "", // external Dockerfile or tarball
		NoCache:        true,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     false,
		Isolation:      "",
		CPUSetCPUs:     "",
		CPUSetMems:     "",
		CPUShares:      0,
		CPUQuota:       0,
		CPUPeriod:      0,
		Memory:         0,
		MemorySwap:     0,
		CgroupParent:   "",
		NetworkMode:    NetworkBridge,
		ShmSize:        0,
		Dockerfile:     "Dockerfile",
		Ulimits:        nil,
		BuildArgs:      nil,
		AuthConfigs: map[string]registry.AuthConfig{
			c.registryAddr: c.GetRegistryAuthConfig(),
		},
		Context:     nil,
		Labels:      nil,
		Squash:      false,
		CacheFrom:   nil,
		SecurityOpt: nil,
		ExtraHosts:  nil,
		Target:      "",
		SessionID:   "",
		Platform:    PlatformLinuxAmd64,
		Version:     "1",
		BuildID:     "",
		Outputs:     nil,
	}
}

const (
	PlatformLinuxAmd64   = internal_api.Linux + "/" + internal_api.Amd64
	PlatformLinuxArm64   = internal_api.Linux + "/" + internal_api.Arm64
	PlatformDarwinArm64  = internal_api.Darwin + "/" + internal_api.Arm64
	PlatformDarwinAmd64  = internal_api.Darwin + "/" + internal_api.Amd64
	PlatformWindowsAmd64 = internal_api.Windows + "/" + internal_api.Amd64
)

const (
	NetworkBridge    = "bridge"
	NetworkHost      = "host"
	NetworkNone      = "none"
	networkContainer = "container:<name|id>"
)
