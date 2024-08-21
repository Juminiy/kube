package kapi

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	"time"
)

var (
	PodStateSyncingDone = make(chan struct{})
)

func SyncPodListByNS(
	dConf *DeploymentConfig,
	qDur time.Duration,
) {
	dConf.CallBack.List = func() error {
		latestDeployListOf := dConf.CallBack.latest.(*v1.DeploymentList)
		for _, deployOf := range latestDeployListOf.Items {
			fmt.Println(deployOf.Name, deployOf.Status, time.Since(deployOf.CreationTimestamp.Time))
		}
		return nil
	}

	go func() {
		defer func() {
			fmt.Println("has return from go func(){}")
		}()
		for {
			select {
			case <-PodStateSyncingDone:
				fmt.Printf("close syncing NS[%s] pods state\n", dConf.Namespace)
				return
			default:
				fmt.Printf("syncing NS[%s] pods state by Dur[%v]\n", dConf.Namespace, qDur)
				err := dConf.List()
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(qDur)
			}
		}
	}()

}
