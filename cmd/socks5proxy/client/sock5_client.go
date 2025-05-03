package main

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/armon/go-socks5"
	"os"
)

func main() {
	_, err := socks5.NewRequest(os.Stdin)
	util.Must(err)

}
