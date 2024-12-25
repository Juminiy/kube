package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_cast"
	"github.com/Juminiy/kube/pkg/util/safe_reflect"
	psmem "github.com/shirou/gopsutil/v4/mem"
)

type mem struct {
	Total     uint64 `json:"total"`
	Available uint64 `json:"available"`
	Used      uint64 `json:"used"`

	TotalSize     string `json:"totalSize"`
	AvailableSize string `json:"availableSize"`
	UsedSize      string `json:"usedSize"`
}

func (m *mem) setHumanRead() *mem {
	m.TotalSize = util.MeasureByte(safe_cast.U64toI(m.Total))
	m.AvailableSize = util.MeasureByte(safe_cast.U64toI(m.Available))
	m.UsedSize = util.MeasureByte(safe_cast.U64toI(m.Used))
	return m
}

func vmem() *mem {
	vmemStat, err := psmem.VirtualMemory()
	if err != nil {
		stdlog.ErrorF("psutil virtual memory error: %s", err.Error())
		return nil
	}
	memPtr := &mem{}
	safe_reflect.CopyFieldValue(vmemStat, memPtr)
	return memPtr.setHumanRead()
}

type MemV2 struct {
	*mem `json:",inline"`
}

func memV2() MemV2 {
	return MemV2{mem: vmem()}
}
