package main

import (
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/file"
	"os/exec"
)

// run in main package, not go-generate
func main() {
	pathList := []string{"pkg", "util", "random", "mock"}
	dirFileNames, err := internal_api.GetDirFileNames(util.GetWorkPath(pathList...))
	util.Must(err)

	for i := range dirFileNames {
		srcPath := util.GetWorkPath(append(pathList, dirFileNames[i])...)
		dstPath := util.GetWorkPath(append(pathList, "v2", dirFileNames[i])...)
		file.ReadGo(srcPath).PackageOf("mockv2").WriteTo(dstPath)
		err = exec.Command("git", "add", dstPath).Start()
		if err != nil {
			stdlog.Error(err)
		}
	}

}
