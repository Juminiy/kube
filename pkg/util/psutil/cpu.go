package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	pscpu "github.com/shirou/gopsutil/cpu"
)

func cpu() []pscpu.InfoStat {
	cpuInfo, err := pscpu.Info()
	if err != nil {
		stdlog.ErrorF("psutil cpu error: %s", err.Error())
		return nil
	}

	return cpuInfo
}
