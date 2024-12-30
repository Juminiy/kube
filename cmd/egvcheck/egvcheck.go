package main

import (
	"fmt"
	"github.com/Juminiy/kube/pkg/util"
	kfile "github.com/Juminiy/kube/pkg/util/file"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"os"
	"path/filepath"
)

// Exported Global Var

func main() {
	readDir(util.GetWorkPath("pkg"))
	displayVar()
}

var pkgs = map[string][]*kfile.ExportedVar{}

func readDir(dirPath string) {
	dir, err := os.ReadDir(dirPath)
	util.Must(err)
	for _, subdir := range dir {
		if subdir.Name() == "testdata" {
			continue
		}
		subDirPath := filepath.Join(dirPath, subdir.Name())
		if subdir.IsDir() {
			readDir(subDirPath)
		} else {
			if varfile := kfile.ReadGoVar(subDirPath); varfile != nil {
				pkgs[dirPath] = append(pkgs[dirPath], varfile)
			}
		}
	}
}

func displayVar() {
	fmt.Println(safe_json.Pretty(pkgs))
}
