package docker_api

import (
	"errors"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_client"
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"io"
)

type BuildImageRespV1 struct {
	docker_client.ImageBuildResp `json:"build_image"`
	TagPushImageResp             `json:"tag_push_image"`
	OSType                       string `json:"os_type"`
	types.ImageBuildResponse     `json:"-"`
}

func (c *Client) BuildImage(input io.Reader, refStr string) (resp BuildImageRespV1, err error) {
	resp.ImageBuildResponse, err = c.buildImage(input, refStr)
	if err != nil {
		return
	}
	resp.parse()
	resp.TagPushImageResp, err = c.tagImageFromRefStr(refStr)
	return
}

func (c *Client) BuildImageFavOption(refStr string) types.ImageBuildOptions {
	return types.ImageBuildOptions{
		Tags:           []string{refStr},
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
		Dockerfile:     "",
		Ulimits:        nil,
		BuildArgs:      nil,
		AuthConfigs: map[string]registry.AuthConfig{
			c.reg.Addr: c.reg.GetAuthConfig(),
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

func (c *Client) buildImage(input io.Reader, refStr string) (types.ImageBuildResponse, error) {
	return c.cli.ImageBuild(c.ctx, input, c.BuildImageFavOption(refStr))
}

var ErrAbsRefStr = errors.New("image absolutely refStr format error")
var ErrProjectNotFound = errors.New("project not found error")

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

func (b *BuildImageRespV1) parse() *BuildImageRespV1 {
	bs, err := io.ReadAll(b.Body)
	defer util.SilentCloseIO("build image resp body", b.Body)
	if err != nil {
		stdlog.ErrorF("read build image bytes error: %s", err.Error())
		return b
	}
	b.ImageBuildResp = (&docker_client.EventResp{}).
		ParseBytes(bs).GetImageBuildResp()
	return b
}
