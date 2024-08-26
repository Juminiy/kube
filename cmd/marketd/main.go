package main

import (
	"fmt"
	ldversion "kube/version"
)

func main() {
	ldversion.Info()

	fmt.Println("marketd://http")
}
