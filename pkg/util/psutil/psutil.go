package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/Juminiy/kube/pkg/util/safe_json"
	psnvgpu "github.com/mindprince/gonvml"
	pscpu "github.com/shirou/gopsutil/v4/cpu"
	psdisk "github.com/shirou/gopsutil/v4/disk"
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
		Nvidia   []psnvgpu.Device `json:"nvidia"` // Nvidia GPU
		Apple    []struct{}       `json:"apple"`  // Apple Silicon MPS
		Amd      []struct{}       `json:"amd"`
		Intel    []struct{}       `json:"intel"`
		Qualcomm []struct{}       `json:"qualcomm"`
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
	bs, err := safe_json.STD().Marshal(getSysHard())
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

func MarshalIndentString() string {
	return util.Bytes2StringNoCopy(MarshalIndent())
}
