package pod_api

import (
	"encoding/json"
	"time"

	"github.com/Juminiy/kube/pkg/log_api/zaplog"
	"github.com/Juminiy/kube/pkg/util"

	appsv1 "k8s.io/api/apps/v1"
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
			statusBytes, err := json.MarshalIndent(deployOf.Status, "", util.JSONMarshalIndent)
			if err != nil {
				zaplog.Error(err)
				return err
			}
			zaplog.Info(deployOf.Name, time.Since(deployOf.CreationTimestamp.Time), util.Bytes2StringNoCopy(statusBytes))
		}
		return nil
	}

	go func() {
		for {
			select {
			case <-PodStateSyncingDone:
				zaplog.InfoF("stop syncing NS[%s] pods state\n", dConf.Namespace)
				return
			default:
				zaplog.InfoF("start syncing NS[%s] pods state by Dur[%v]\n", dConf.Namespace, qDur)
				err := dConf.List()
				if err != nil {
					zaplog.Error(err)
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
