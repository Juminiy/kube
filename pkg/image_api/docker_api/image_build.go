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
	"github.com/samber/lo"
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
		Version:     types.BuilderBuildKit, // use BuildKit
		BuildID:     "",
		Outputs:     nil,
	}
}

type BuildImageRespV2 struct {
	docker_client.ImageBuildResp `json:"build_image"`
}

func (c *Client) BuildImageV3(ctx context.Context, input io.Reader, options types.ImageBuildOptions) (
	resp BuildImageRespV2, err error) {
	c.supplementImageBuildOptions(&options)
	buildResp, err := c.apiClient.ImageBuild(input, options, ctx)
	if err != nil {
		return
	}
	resp.ImageBuildResp = buildResp.GetImageBuildResp()
	return
}

func (c *Client) supplementImageBuildOptions(options *types.ImageBuildOptions) {
	validAbsRefStr := make([]string, 0, len(options.Tags))
	for i, refStr := range options.Tags {
		arti := ParseToArtifact(refStr)
		if len(arti.Registry) == 0 {
			arti.SetRegistry(c.reg.Addr)
		}
		if len(arti.Project) == 0 {
			arti.SetProject(c.reg.GetProject())
		}
		options.Tags[i] = arti.AbsRefStr()
		if arti.ValidAbsRefStr() {
			validAbsRefStr = append(validAbsRefStr, arti.AbsRefStr())
		}
	}

	options.SuppressOutput = false
	options.Version = types.BuilderBuildKit

	if len(options.AuthConfigs) == 0 {
		options.AuthConfigs = map[string]registry.AuthConfig{
			c.reg.Addr: c.reg.GetAuthConfig(),
		}
	}

	if _, ok := lo.Find(options.Outputs, func(item types.ImageBuildOutput) bool {
		if item.Type == OutputRegistry {
			return true
		}
		return false
	}); ok {
		options.Outputs = make([]types.ImageBuildOutput, len(validAbsRefStr))
		for i := range validAbsRefStr {
			options.Outputs[i] = types.ImageBuildOutput{
				Type: OutputImage,
				Attrs: map[string]string{
					"name": validAbsRefStr[i],
					"push": "true",
				},
			}
		}
	}

}

var ErrAbsRefStr = errors.New("image absolutely refStr format error")
var ErrProjectNotFound = errors.New("project not found error")

const (
	/* os[/arch[/variant]] */
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
	networkContainer = "container:<name|id>" // format only, not to use
)

const (
	OutputLocal    = "local"
	OutputTar      = "tar"
	OutputOCI      = "oci"
	OutputDocker   = "docker"
	OutputImage    = "image"
	OutputRegistry = "registry"
)

type BuildOutput struct {
	Local  []string // destination directory where files will be written
	Tar    []string // destination path where tarball will be written. “-” writes to stdout.
	OCI    []string // destination path where tarball will be written. “-” writes to stdout.
	Docker []struct {
		Dest    string // destination path where tarball will be written. If not specified, the tar will be loaded automatically to the local image store.
		Context string // name for the Docker context where to import the result
	}
	Image []struct {
		Name string // name (references) for the new image.
		Push bool   // Boolean to automatically push the image.
	}
	Registry bool // The registry exporter is a shortcut for type=image,push=true.
}
