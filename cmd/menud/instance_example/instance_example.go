package instance_example

import (
	"fmt"
	"kube/pkg/instance_example/nginx"
	"kube/pkg/instance_example/ubuntu"
	"kube/pkg/k8s_api"
	"os"
	"time"
)

func Menu(s ...string) {
	var (
		dConf    *k8s_api.DeploymentConfig
		appOf    string
		actionOf string

		err error
	)
	switch len(s) {
	case 1:
		appOf = s[0]
	case 2:
		appOf, actionOf = s[0], s[1]
	default:
		return
	}

	switch appOf {
	case "nginx":
		dConf = nginx.NewNginxDeployment()
	case "ubuntu":
		dConf = ubuntu.NewUbuntuDeployment()
	//case "centos":
	//case "tensorflow":
	//case "pytorch":
	//case "LLM IDE":
	default:
		fmt.Printf("error app: %v\n", appOf)
		return
	}

	switch actionOf {
	case "create", "c":
		fmt.Println(appOf, "deployment creating...")
		err = dConf.Create()
	case "update", "u":
		fmt.Println(appOf, "deployment updating...")
		err = dConf.Update()
	case "delete", "d":
		fmt.Println(appOf, "deployment deleting...")
		err = dConf.Delete()
	case "stop", "s":
		fmt.Println(appOf, "deployment stopping...")
		var saveCfg string
		saveCfg, err = dConf.Stop()
		fmt.Println(saveCfg)
	case "list", "l":
		fmt.Println("list deployment in NS default")
		err = dConf.List()
	case "quit", "q":
		fmt.Println("quit command line interaction interface")
		os.Exit(0)
	case "start-sync":
		k8s_api.SyncPodListByNS(dConf, 5*time.Second)
	case "stop-sync":
		fmt.Println("stop the pod state syncing...")
		k8s_api.PodStateSyncingDone <- struct{}{}
	default:
		fmt.Printf("error action: %v\n", actionOf)
	}
	if err != nil {
		fmt.Printf("unknown error: %v\n", err)
	}
}
