package pod_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	"time"

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
	dConf.cbk.List = func() error {
		latestDeployListOf := dConf.cbk.Latest.(*appsv1.DeploymentList)
		if latestDeployListOf == nil ||
			len(latestDeployListOf.Items) == 0 {
			PodStateSyncingDone <- struct{}{}
		}
		for _, deployOf := range latestDeployListOf.Items {
			statusBytes, err := safe_json.STD().MarshalIndent(deployOf.Status, "", util.JSONMarshalIndent)
			if err != nil {
				stdlog.Error(err)
				return err
			}
			stdlog.Info(deployOf.Name, time.Since(deployOf.CreationTimestamp.Time), util.Bytes2StringNoCopy(statusBytes))
		}
		return nil
	}

	go func() {
		for {
			select {
			case <-PodStateSyncingDone:
				stdlog.InfoF("stop syncing NS[%s] pods state\n", dConf.Namespace)
				return
			default:
				stdlog.InfoF("start syncing NS[%s] pods state by Dur[%v]\n", dConf.Namespace, qDur)
				err := dConf.List()
				if err != nil {
					stdlog.Error(err)
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
