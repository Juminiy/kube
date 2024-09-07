package main

import (
	"flag"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	ldversion "github.com/Juminiy/kube/version"
)

func main() {
	initFlag()
	ldversion.Info(version)

	stdlog.InfoW("http service", "name", "marketd")
}

// global Flags
var (
	version *bool
)

func initFlag() {
	version = flag.Bool("v", false, "print version json info")
	flag.Parse()
}
