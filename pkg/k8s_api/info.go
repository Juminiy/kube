package k8s_api

import (
	"encoding/json"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	"kube/pkg/util"
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
		latestDeployListOf := dConf.CallBack.latest.(*appsv1.DeploymentList)
		if latestDeployListOf == nil ||
			len(latestDeployListOf.Items) == 0 {
			PodStateSyncingDone <- struct{}{}
		}
		for _, deployOf := range latestDeployListOf.Items {
			status, err := json.MarshalIndent(deployOf.Status, "", util.JSONMarshalIndent)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(deployOf.Name, time.Since(deployOf.CreationTimestamp.Time), string(status))
		}
		return nil
	}

	go func() {
		for {
			select {
			case <-PodStateSyncingDone:
				fmt.Printf("stop syncing NS[%s] pods state\n", dConf.Namespace)
				return
			default:
				fmt.Printf("start syncing NS[%s] pods state by Dur[%v]\n", dConf.Namespace, qDur)
				err := dConf.List()
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(qDur)
			}
		}
	}()

}

func ListNodes() {

}

func ListPods() {

}

func ListDeployments() {

}
