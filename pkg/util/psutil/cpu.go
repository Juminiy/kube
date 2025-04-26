package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	pscpu "github.com/shirou/gopsutil/v4/cpu"
)

func cpu() []pscpu.InfoStat {
	cpuInfo, err := pscpu.Info()
	if err != nil {
		stdlog.ErrorF("psutil cpu error: %s", err.Error())
		return nil
	}

	return cpuInfo
}

type CPUV2 struct {
	ModelCount int
	CPUModel   map[string]CPUByModel
}

type CPUByModel struct {
	Cores         int32
	CacheSizeByte int32
	CacheSize     string
	Frequency     string
	Flags         []string
	InfoStat      []pscpu.InfoStat `json:"cpu_raw_info_stat"`
}

func cpuV2() CPUV2 {
	cpuInfo := cpu()
	cpuModel := make(map[string]CPUByModel, util.MagicMapCap)
	for _, info := range cpuInfo {
		modelInfo := cpuModel[info.ModelName]
		modelInfo.Cores += info.Cores
		modelInfo.CacheSizeByte += info.CacheSize
		modelInfo.CacheSize = util.MeasureByte(int(modelInfo.CacheSizeByte))
		modelInfo.Frequency = util.F64toa(info.Mhz) + "MHZ"
		modelInfo.Flags = info.Flags
		info.Flags = nil
		modelInfo.InfoStat = append(modelInfo.InfoStat, info)
		cpuModel[info.ModelName] = modelInfo
	}
	return CPUV2{
		ModelCount: len(cpuModel),
		CPUModel:   cpuModel,
	}
}
