package version

import (
	"flag"
	"fmt"
	"os"
)

var info struct {
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

func Info() {
	if len(os.Args) == 2 {
		version := flag.Bool("v", false, "print version information and exit")
		flag.Parse()
		if *version {
			fmt.Println("Kube version:", info)
			os.Exit(0)
		}
	}
}
