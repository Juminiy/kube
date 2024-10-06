package psutil

import (
	internaldev "github.com/Juminiy/kube/pkg/internal/device"
	"github.com/Juminiy/kube/pkg/internal_api"
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	psnvgpu "github.com/mindprince/gonvml"
	"runtime"
)

func gpu() *sysGPU {
	sysgpu := &sysGPU{}
	switch {
	case runtime.GOOS == internal_api.Darwin && runtime.GOARCH == internal_api.Arm64:
		sysgpu.Apple = internaldev.MPS()
	default:
		sysgpu.Nvidia = nvidiaGPU()
	}
	return sysgpu
}

func nvidiaGPU() []psnvgpu.Device {
	devCnt, err := psnvgpu.DeviceCount()
	if err != nil {
		stdlog.ErrorF("psutil nvidia gpu count error: %s", err.Error())
		return nil
	}
	gpuDevices := make([]psnvgpu.Device, devCnt)
	for devID := range devCnt {
		gpuDevices[devID], err = psnvgpu.DeviceHandleByIndex(devID)
		if err != nil {
			stdlog.ErrorF("psutil nvidia gpu devID: %d error: %s", devID, err.Error())
			return nil
		}
	}
	return gpuDevices
}
