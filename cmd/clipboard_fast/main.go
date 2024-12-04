package main

import (
	"github.com/Juminiy/kube/pkg/netserver/http/fiberserver"
)

func main() {
	fiberserver.New().WithPort(8081).Load()
}
