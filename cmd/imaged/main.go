package main

import (
	"fmt"
	ldversion "github.com/Juminiy/kube/version"
)

func main() {
	ldversion.Info()

	fmt.Println("imaged://http")
}
