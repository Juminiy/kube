package main

import (
	"fmt"
	"kube/deploy_example"
	"kube/harbor_api"
	"kube/k8s_api"
	ldversion "kube/version"
)

func main() {
	ldversion.Info()

	var (
		moduleOf string
		appOf    string
		actionOf string
	)
	for {
		fmt.Println("module | app | action")
		fmt.Println("module: cluster | deploy | harbor")
		fmt.Println("app: [log | node | deployment] | [nginx | ubuntu]")
		fmt.Println("action: [create | update | delete | list | start-sync | stop-sync]")
		fmt.Println("Ctrl+C: exit(0) to exit")
		if _, err := fmt.Scanf("%s %s %s", &moduleOf, &appOf, &actionOf); err != nil {
			fmt.Printf("error input: %v\n", err)
		}
		switch moduleOf {
		case "cluster":
			k8s_api.Menu(appOf, actionOf)
		case "deploy":
			deploy_example.Menu(appOf, actionOf)
		case "harbor":
			harbor_api.Menu()
		}
	}

}
