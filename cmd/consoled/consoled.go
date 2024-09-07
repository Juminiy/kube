package main

import (
	"flag"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	ldversion "github.com/Juminiy/kube/version"
)

func main() {
	initFlag()
	ldversion.Info(version)
	if developer != nil {
		stdlog.Info("developer:", *developer)
	}

	stdlog.InfoW("http service", "name", "consoled")
}

// global Flags
var (
	version   *bool
	developer *string
)

func initFlag() {
	version = flag.Bool("v", false, "print version json info")
	developer = flag.String("d", "chisato", "print developer of consoled")
	flag.Parse()
}
