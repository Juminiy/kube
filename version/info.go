package version

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"os"
	"runtime"
)

type (
	infoType struct {
		GitMajor     string
		GitMinor     string
		GitVersion   string
		GitCommit    string
		GitTreeState string
		BuildDate    string
		GoVersion    string
		Compiler     string
		Platform     string
	}
)

var (
	GitMajor     string // major version, always numeric
	GitMinor     string // minor version, numeric possibly followed by "+"
	GitVersion   string // semantic version, derived by build scripts
	GitCommit    string // sha1 from git, output of $(git rev-parse HEAD)
	GitTreeState string // state of git tree, either "clean" or "dirty"
	BuildDate    string // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	info         = infoType{
		GitMajor:     GitMajor,
		GitMinor:     GitMinor,
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
)

func Info() {
	if len(os.Args) == 2 {
		version := flag.Bool("v", false, "print version json info")
		flag.Parse()
		if *version {
			infoJSON, err := util.MarshalJSONPretty(&info)
			if err != nil {
				fmt.Println(err)
				goto osExit
			}
			fmt.Println(infoJSON)
		}
	osExit:
		os.Exit(0)
	}
}
