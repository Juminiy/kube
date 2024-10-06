package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	psdisk "github.com/shirou/gopsutil/disk"
)

func disk() []psdisk.PartitionStat {
	diskParts, err := psdisk.Partitions(true)
	if err != nil {
		stdlog.ErrorF("psutil disk error: %s", err.Error())
		return nil
	}
	return diskParts
}
