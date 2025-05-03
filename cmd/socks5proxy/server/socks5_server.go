package main

import (
	"github.com/Juminiy/kube/pkg/util"
	"github.com/armon/go-socks5"
)

func main() {
	srv, err := socks5.New(&socks5.Config{})
	util.Must(err)

	util.Must(srv.ListenAndServe("tcp", "127.0.0.1:7651"))
}
