package version

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"runtime"
)

type (
	infoType struct {
		Major        string `json:"major,omitempty"`
		Minor        string `json:"minor,omitempty"`
		GitVersion   string `json:"gitVersion,omitempty"`
		GitCommit    string `json:"gitCommit,omitempty"`
		GitTreeState string `json:"gitTreeState,omitempty"`
		BuildDate    string `json:"buildDate,omitempty"`
		GoVersion    string `json:"goVersion,omitempty"`
		Compiler     string `json:"compiler,omitempty"`
		Platform     string `json:"platform,omitempty"`
	}
)

var (
	GitMajor     string // major version, always numeric
	GitMinor     string // minor version, numeric possibly followed by "+"
	GitVersion   string // semantic version, derived by build scripts
	GitCommit    string // sha1 from git, output of $(git rev-parse HEAD)
	GitTreeState string // state of git tree, either "clean" or "dirty"
	BuildDate    string // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	info         infoType
)

func init() {
	flag.StringVar(&GitMajor, "GitMajor", "", "")
	flag.StringVar(&GitMinor, "GitMinor", "", "")
	flag.StringVar(&GitVersion, "GitVersion", "", "")
	flag.StringVar(&GitCommit, "GitCommit", "", "")
	flag.StringVar(&GitTreeState, "GitTreeState", "", "")
	flag.StringVar(&BuildDate, "BuildDate", "", "")

	info = infoType{
		Major:        GitMajor,
		Minor:        GitMinor,
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func Info() {
	if len(os.Args) == 2 {
		version := flag.Bool("v", false, "print version info, -v")
		infoFormat := flag.String("f", "json", "print version format, -f json | yaml | string")
		flag.Parse()

		if *version {
			if infoFormat != nil {
				switch *infoFormat {
				case "j", "json":
					j, err := json.MarshalIndent(&info, "", "")
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(string(j))
				case "y", "yaml":
					y, err := yaml.Marshal(&info)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(string(y))
				default:
					fmt.Printf("Kube version: %#v\n", info)
				}
			} else {
				fmt.Printf("Kube version: %#v\n", info)
			}
			os.Exit(0)
		}
	}
}
