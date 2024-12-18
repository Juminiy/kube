package cmd

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	buildkitclient "github.com/moby/buildkit/client"
	"github.com/samber/lo"
)

type cmder struct{}

var _cmder = cmder{}

func Get() cmder { return _cmder }

func (c cmder) Cmd(options types.ImageBuildOptions) []string {
	if options.Version == types.BuilderBuildKit {
		return c.buildKit(options)
	}
	return []string{"buildctl", "-v"}
}

func (cmder) buildKit(options types.ImageBuildOptions) []string {
	cmd := make([]string, 0, util.MagicSliceCap<<2)
	cmd = append(cmd,
		"buildctl", "build",
		"--frontend", "dockerfile.v0",
		"--progress", "rawjson",
		"--local", "context=.",
		"--local", "dockerfile=.",
		"--debug-json-cache-metrics", "stdout")

	imageOutputs := lo.FilterMap(options.Outputs, func(item types.ImageBuildOutput, index int) (types.ImageBuildOutput, bool) {
		if item.Type == buildkitclient.ExporterImage || item.Type == BuildExporterRegistry {
			return item, true
		}
		return item, false
	})
	if len(imageOutputs) > 0 && len(options.Tags) > 0 {
		for _, refStr := range options.Tags {
			cmd = append(cmd,
				"--output", fmt.Sprintf("type=image,name=%s,push=true", refStr))
		}
	}

	if options.NoCache {
		cmd = append(cmd, "--no-cache")
	}

	for buildArgk, buildArgv := range options.BuildArgs {
		if buildArgv != nil {
			cmd = append(cmd,
				"--opt", fmt.Sprintf("build-arg:%s=%s", buildArgk, *buildArgv))
		}
	}

	if len(options.Target) > 0 {
		cmd = append(cmd,
			"--opt", fmt.Sprintf("target=%s", options.Target))
	}

	return cmd
}

const (
	BuildExporterRegistry = "registry"
)
