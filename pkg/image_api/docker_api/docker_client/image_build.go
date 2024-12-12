package docker_client

import (
	"context"
	"encoding/base64"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"io"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) ImageBuild(input io.Reader, buildOptions types.ImageBuildOptions, ctx ...context.Context) (
	buildResp EventResp, err error) {
	query, err := c.imageBuildOptionsToQuery(c.ctx, buildOptions)
	if err != nil {
		return
	}

	authBuf, err := EncE(buildOptions.AuthConfigs)
	if err != nil {
		return
	}

	r := c.post("/build")
	r.r.SetHeader("X-Registry-Config", base64.URLEncoding.EncodeToString(authBuf)).
		SetHeader("Content-Type", "application/x-tar").
		SetQueryParamsFromValues(query).
		SetBody(input)

	resp, err := r.do()
	return *buildResp.Parse(resp), err
}

// copy from
func (c *Client) imageBuildOptionsToQuery(ctx context.Context, options types.ImageBuildOptions) (url.Values, error) {
	query := url.Values{
		"t":           options.Tags,
		"securityopt": options.SecurityOpt,
		"extrahosts":  options.ExtraHosts,
	}
	if options.SuppressOutput {
		query.Set("q", "1")
	}
	if options.RemoteContext != "" {
		query.Set("remote", options.RemoteContext)
	}
	if options.NoCache {
		query.Set("nocache", "1")
	}
	if options.Remove {
		query.Set("rm", "1")
	} else {
		query.Set("rm", "0")
	}

	if options.ForceRemove {
		query.Set("forcerm", "1")
	}

	if options.PullParent {
		query.Set("pull", "1")
	}

	if options.Squash {
		//if err := cli.NewVersionError(ctx, "1.25", "squash"); err != nil {
		//	return query, err
		//}
		query.Set("squash", "1")
	}

	if !container.Isolation.IsDefault(options.Isolation) {
		query.Set("isolation", string(options.Isolation))
	}

	query.Set("cpusetcpus", options.CPUSetCPUs)
	query.Set("networkmode", options.NetworkMode)
	query.Set("cpusetmems", options.CPUSetMems)
	query.Set("cpushares", strconv.FormatInt(options.CPUShares, 10))
	query.Set("cpuquota", strconv.FormatInt(options.CPUQuota, 10))
	query.Set("cpuperiod", strconv.FormatInt(options.CPUPeriod, 10))
	query.Set("memory", strconv.FormatInt(options.Memory, 10))
	query.Set("memswap", strconv.FormatInt(options.MemorySwap, 10))
	query.Set("cgroupparent", options.CgroupParent)
	query.Set("shmsize", strconv.FormatInt(options.ShmSize, 10))
	query.Set("dockerfile", options.Dockerfile)
	query.Set("target", options.Target)

	ulimitsJSON, err := EncE(options.Ulimits)
	if err != nil {
		return query, err
	}
	query.Set("ulimits", string(ulimitsJSON))

	buildArgsJSON, err := EncE(options.BuildArgs)
	if err != nil {
		return query, err
	}
	query.Set("buildargs", string(buildArgsJSON))

	labelsJSON, err := EncE(options.Labels)
	if err != nil {
		return query, err
	}
	query.Set("labels", string(labelsJSON))

	cacheFromJSON, err := EncE(options.CacheFrom)
	if err != nil {
		return query, err
	}
	query.Set("cachefrom", string(cacheFromJSON))
	if options.SessionID != "" {
		query.Set("session", options.SessionID)
	}
	if options.Platform != "" {
		//if err := cli.NewVersionError(ctx, "1.32", "platform"); err != nil {
		//	return query, err
		//}
		query.Set("platform", strings.ToLower(options.Platform))
	}
	if options.BuildID != "" {
		query.Set("buildid", options.BuildID)
	}
	query.Set("version", string(options.Version))

	if options.Outputs != nil {
		outputsJSON, err := EncE(options.Outputs)
		if err != nil {
			return query, err
		}
		query.Set("outputs", string(outputsJSON))
	}
	return query, nil
}

type ImageBuildResp struct {
	Streams  []string `json:"image_build_info"`
	Sha256ID string   `json:"image_sha256_id"`
}

func (r *EventResp) GetImageBuildResp() (resp ImageBuildResp) {
	resp.Streams = make([]string, 0, len(r.Message))
	for _, msg := range r.Message {
		if msg == nil {
			return
		}
		if len(msg.Stream) != 0 {
			resp.Streams = append(resp.Streams, msg.Stream)
		}
		if msg.Aux != nil {
			var auxID = struct {
				ID string
			}{}
			err := DecE(*msg.Aux, &auxID)
			if err != nil {
				stdlog.ErrorF("unmarshal auxID error: %s", err.Error())
			}
		}
	}
	return
}
