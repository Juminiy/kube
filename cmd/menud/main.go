package main

import (
	"fmt"
	ldversion "github.com/Juminiy/kube/version"
	"os"

	harbormenu "github.com/Juminiy/kube/cmd/menud/harbor_api"
	instancemenu "github.com/Juminiy/kube/cmd/menud/instance_example"
	k8smenu "github.com/Juminiy/kube/cmd/menud/k8s_api"
)

func main() {
	ldversion.Info()

	var (
		setting                   string
		moduleOf, appOf, actionOf string
	)
	const (
		helpRetCode int8 = 0
		nextRetCode int8 = 1
	)

	helpMenu := func(s ...string) int8 {
		if s[0] == "help" || s[0] == "h" {
			fmt.Println("help | quit | none [module | app | action]")
			fmt.Println("help: help | h")
			fmt.Println("quit: quit | q")
			fmt.Println("next: next | n")
			fmt.Println("module: cluster | deploy | harbor")
			fmt.Println("app: [log | node | deployment] | [nginx | ubuntu]")
			fmt.Println("action: [create | update | delete | list | start-sync | stop-sync]")
			return helpRetCode
		} else if s[0] == "quit" || s[0] == "q" {
			os.Exit(0)
		}
		return nextRetCode
	}
	for {
		fmt.Printf("setting [help | quit | next]: ")
		if _, err := fmt.Scanf("%s", &setting); err != nil {
			fmt.Printf("error setting: %v\n", err)
		}
		if helpMenu(setting) == helpRetCode {
			continue
		}
		fmt.Printf("module [cluster | deploy | harbor]: ")
		if _, err := fmt.Scanf("%s %s %s", &moduleOf, &appOf, &actionOf); err != nil {
			fmt.Printf("error input: %v\n", err)
			return
		}
		switch moduleOf {
		case "cluster":
			k8smenu.Menu(appOf, actionOf)
		case "deploy":
			instancemenu.Menu(appOf, actionOf)
		case "harbor":
			harbormenu.Menu(appOf, actionOf)
		}
	}

}
