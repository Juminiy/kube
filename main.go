package main

import (
	"fmt"
	"kube/deploy_example"
	"kube/kapi"
	ldversion "kube/version"
	"time"
)

func main() {
	ldversion.Info()

	var (
		dConf  = deploy_example.NewNginxDeployment()
		action string
		err    error
	)
	for {
		if _, err = fmt.Scanf("%s", &action); err != nil {
			fmt.Printf("error input: %v\n", err)
			return
		}
		switch action {
		case "create", "c":
			fmt.Println("nginx deployment creating...")
			err = dConf.Create()
		case "update", "u":
			fmt.Println("nginx deployment updating...")
			err = dConf.Update()
		case "delete", "d":
			fmt.Println("nginx deployment deleting...")
			err = dConf.Delete()
		case "list", "l":
			fmt.Println("list deployment in NS default")
			err = dConf.List()
		case "quit", "q":
			fmt.Println("quit command line interaction interface")
			return
		case "start-sync":
			kapi.SyncPodListByNS(dConf, 5*time.Second)
		case "stop-sync":
			fmt.Println("stop the pod state syncing...")
			kapi.PodStateSyncingDone <- struct{}{}
		default:
			fmt.Printf("error action: %v\n", action)
		}
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

}
