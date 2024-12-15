package version

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Juminiy/kube/pkg/util"
)

var (
	GitMajor     string // major version, always numeric
	GitMinor     string // minor version, numeric possibly followed by "+"
	GitPatch     string // patch version, always numeric
	GitVersion   string // semantic version, derived by build scripts
	GitCommit    string // sha1 from git, output of $(git rev-parse HEAD)
	GitTreeState string // state of git tree, either "clean" or "dirty"
	BuildDate    string // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

var info = infoType{
	GitMajor:     GitMajor,
	GitMinor:     GitMinor,
	GitPatch:     GitPatch,
	GitVersion:   GitVersion,
	GitCommit:    GitCommit,
	GitTreeState: GitTreeState,
	BuildDate:    BuildDate,
	GoVersion:    runtime.Version(),
	Compiler:     runtime.Compiler,
	Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
}

type infoType struct {
	GitMajor     string
	GitMinor     string
	GitPatch     string
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

func Info(version *bool) {
	if version != nil && *version {
		infoJSON, err := util.MarshalJSONPretty(&info)
		if err != nil {
			util.DevPanic(err)
		}
		fmt.Println(infoJSON)
		os.Exit(0)
	}
}
