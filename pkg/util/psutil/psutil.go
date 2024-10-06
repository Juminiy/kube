package psutil

import (
	"encoding/json"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	psnvgpu "github.com/mindprince/gonvml"
	pscpu "github.com/shirou/gopsutil/cpu"
	psdisk "github.com/shirou/gopsutil/disk"
	"net"
	"runtime"
)

type (
	sysHard struct {
		OS          string                 `json:"os"`
		Arch        string                 `json:"arch"`
		Mem         *mem                   `json:"mem"`
		Disk        []psdisk.PartitionStat `json:"disk"`
		Network     []net.Interface        `json:"network"`
		CPU         []pscpu.InfoStat       `json:"cpu"`
		GPU         *sysGPU                `json:"gpu"`
		MotherBoard any                    `json:"board"`
	}

	sysGPU struct {
		Nvidia []psnvgpu.Device `json:"nvidia"`
	}
)

func getSysHard() *sysHard {
	return &sysHard{
		OS:      runtime.GOOS,
		Arch:    runtime.GOARCH,
		Mem:     vmem(),
		Disk:    disk(),
		Network: netInterface(),
		CPU:     cpu(),
		GPU:     gpu(),
	}
}

func Marshal() []byte {
	bs, err := json.Marshal(getSysHard())
	if err != nil {
		stdlog.ErrorF("marshal psutil error: %s", err.Error())
		return nil
	}
	return bs
}

func MarshalIndent() []byte {
	bs, err := util.MarshalJSONPretty(getSysHard())
	if err != nil {
		stdlog.ErrorF("marshal indent psutil error: %s", err.Error())
		return nil
	}
	return util.String2BytesNoCopy(bs)
}
