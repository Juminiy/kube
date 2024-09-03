package main

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	ldversion "github.com/Juminiy/kube/version"
)

func main() {
	ldversion.Info()

	stdlog.InfoW("http service", "name", "payd")
}
