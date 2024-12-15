package docker_api

import (
	"context"
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
}

func (b *BuildImageRespV1) parseBuildRawResp(brr types.ImageBuildResponse) *BuildImageRespV1 {
	bs, err := io.ReadAll(brr.Body)
	defer util.SilentCloseIO("build image resp body", brr.Body)
	if err != nil {
		stdlog.ErrorF("read build image bytes error: %s", err.Error())
		return b
	}
	b.ImageBuildResp = (&docker_client.EventResp{}).
		ParseBytes(bs).GetImageBuildResp()
	b.OSType = brr.OSType
	return b
}

func (c *Client) BuildImage(input io.Reader, refStr string) (
	resp BuildImageRespV1, err error) {
	return c.buildImageWithContext(c.ctx, input, refStr)
}

func (c *Client) BuildImageWithCancel(ctx context.Context, input io.Reader, refStr string) (
	resp BuildImageRespV1, cancelFunc *context.CancelFunc, err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cancelFunc = &cancel
	resp, err = c.buildImageWithContext(ctx, input, refStr)
	return
}

func (c *Client) buildImageWithContext(ctx context.Context, input io.Reader, refStr string) (
	resp BuildImageRespV1, err error) {
	rawBuildResp, err := c.cli.ImageBuild(ctx, input, c.BuildImageFavOption(refStr))
	if err != nil {
		return
	}
	resp.parseBuildRawResp(rawBuildResp)
	resp.TagPushImageResp, err = c.tagImageFromRefStr(refStr)
	return
}

func (c *Client) BuildImageV2(input io.Reader, refStr string) (
	resp BuildImageRespV1, err error) {
	return c.buildImageWithContextV2(c.ctx, input, refStr)
}

func (c *Client) buildImageWithContextV2(ctx context.Context, input io.Reader, refStr string) (
	resp BuildImageRespV1, err error) {
	buildResp, err := c.apiClient.ImageBuild(input, c.BuildImageFavOption(refStr), ctx)
	if err != nil {
		return
	}
	resp.ImageBuildResp = buildResp.GetImageBuildResp()
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
		Version:     "2", // use BuildKit
		BuildID:     "",
		Outputs:     nil,
	}
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
