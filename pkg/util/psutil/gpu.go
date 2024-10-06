//go:build !(arm64 && darwin)

package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	psnvgpu "github.com/mindprince/gonvml"
)

func gpu() *sysGPU {
	return &sysGPU{Nvidia: nvidiaGPU()}
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
