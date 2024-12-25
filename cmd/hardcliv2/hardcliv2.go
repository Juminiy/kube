package main

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/psutil"
	"github.com/Juminiy/kube/pkg/util/safe_json"
)

func main() {
	ver := flag.Int("v", 2, "psutil version 1 | 2")
	flag.Parse()

	var prettyStr string
	if *ver == 1 {
		prettyStr = util.GreenAny(psutil.MarshalIndentString())
	} else if *ver == 2 {
		prettyStr = util.GreenAny(safe_json.Pretty(psutil.GetSysHardV2()))
	} else {
		prettyStr = util.RedAny(safe_json.Pretty(struct {
			VersionError string `json:"version_error"`
		}{
			VersionError: "hardcli version must 1 or 2",
		}))
	}
	fmt.Println(prettyStr)
}
