package main

import (
	"flag"
	"fmt"
	"github.com/Juminiy/kube/pkg/netserver/http/fiberserver"
	"github.com/gofiber/fiber/v3/log"
)

func main() {}

var (
	_host string
	_port int
)

func init() {
	flag.StringVar(&_host, "host", "0.0.0.0", "net interface to listen on")
	flag.IntVar(&_port, "port", 7788, "port to listen on")
	flag.Parse()
	log.Fatal(fiberserver.Minimal().
		Listen(fmt.Sprintf("%s:%d", _host, _port)))
}
