package main

import udplib "github.com/Juminiy/kube/pkg/netserver/udp"

func main() {
	udplib.IPv4Client("host.local:3344")
}
