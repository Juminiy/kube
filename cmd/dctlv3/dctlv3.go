package main

import (
	"github.com/Juminiy/kube/pkg/util"
	cli "github.com/urfave/cli/v3"
	"golang.org/x/net/context"
	"os"
)

func main() {
	util.Must((&cli.Command{
		Name: "dctlv3",
	}).Run(context.TODO(), os.Args))
}
