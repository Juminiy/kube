package psutil

import (
	"runtime"
)

type SysHard struct {
	OS     string
	Arch   string
	CPU    CPUV2
	Memory MemV2
	Disk   DiskV2
	GPU    GPUV2
	Net    NetV2
}

type GPUV2 struct {
	GPU *sysGPU `json:",inline"`
}

func gpuV2() GPUV2 {
	return GPUV2{GPU: gpu()}
}

func GetSysHardV2() SysHard {
	return SysHard{
		OS:     runtime.GOOS,
		Arch:   runtime.GOARCH,
		CPU:    cpuV2(),
		Memory: memV2(),
		Disk:   diskV2(),
		Net:    netV2(),
		GPU:    gpuV2(),
	}
}
