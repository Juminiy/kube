package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	kubereflect "github.com/Juminiy/kube/pkg/util/reflect"
	psmem "github.com/shirou/gopsutil/v3/mem"
)

type mem struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`
}

func vmem() *mem {
	vmemStat, err := psmem.VirtualMemory()
	if err != nil {
		stdlog.ErrorF("psutil virtual memory error: %s", err.Error())
		return nil
	}
	memPtr := &mem{}
	kubereflect.CopyFieldValue(vmemStat, memPtr)
	return memPtr
}
