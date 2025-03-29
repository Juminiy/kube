package main

import (
	udplib "github.com/Juminiy/kube/pkg/netserver/udp"
)

func main() {
	udplib.IPv4Server("host.local:3344")
}
