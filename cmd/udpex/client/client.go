package main

import udplib "github.com/Juminiy/kube/pkg/netserver/udp"

func main() {
	udplib.IPv4Client("0.0.0.0:3344")
}
