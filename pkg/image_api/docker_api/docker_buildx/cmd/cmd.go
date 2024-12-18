package cmd

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/docker/docker/api/types"
	buildkitclient "github.com/moby/buildkit/client"
	"github.com/samber/lo"
)

/*To build and push an image using Dockerfile:
buildctl build \
	--frontend dockerfile.v0 \
	--opt target=foo \
	--opt build-arg:foo=bar \
	--local context=. \
	--local dockerfile=. \
	--output type=image,name=docker.io/username/image,push=true
*/

/*
OPTIONS:
   --output value, -o value          Define exports for build result, e.g. --output type=image,name=docker.io/username/image,push=true
   --progress value                  Set type of progress (auto, plain, tty, rawjson). Use plain to show container output (default: "auto")
   --trace value                     Path to trace file. Defaults to no tracing.
   --local value                     Allow build access to the local directory
   --oci-layout value                Allow build access to the local OCI layout
   --frontend value                  Define frontend used for build
   --opt value                       Define custom options for frontend, e.g. --opt target=foo --opt build-arg:foo=bar
   --no-cache                        Disable cache for all the vertices
   --export-cache value              Export build cache, e.g. --export-cache type=registry,ref=example.com/foo/bar, or --export-cache type=local,dest=path/to/dir
   --import-cache value              Import build cache, e.g. --import-cache type=registry,ref=example.com/foo/bar, or --import-cache type=local,src=path/to/dir
   --secret value                    Secret value exposed to the build. Format id=secretname,src=filepath
   --allow value                     Allow extra privileged entitlement, e.g. network.host, security.insecure
   --ssh value                       Allow forwarding SSH agent to the builder. Format default|<id>[=<socket>|<key>[,<key>]]
   --metadata-file value             Output build metadata (e.g., image digest) to a file as JSON
   --source-policy-file value        Read source policy file from a JSON file
   --ref-file value                  Write build ref to a file
   --registry-auth-tlscontext value  Overwrite TLS configuration when authenticating with registries, e.g. --registry-auth-tlscontext host=https://myserver:2376,insecure=false,ca=/path/to/my/ca.crt,cert=/path/to/my/cert.crt,key=/path/to/my/key.crt
   --debug-json-cache-metrics value  Where to output json cache metrics, use 'stdout' or 'stderr' for standard (error) output.
*/

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
