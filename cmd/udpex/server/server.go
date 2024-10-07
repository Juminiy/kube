package main

import (
	udplib "github.com/Juminiy/kube/pkg/netserver/udp"
)

func main() {
	udplib.IPv4Server("192.168.3.37:3344")
}
