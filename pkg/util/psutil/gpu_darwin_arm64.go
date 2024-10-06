//go:build darwin && arm64

package psutil

import (
	internaldev "github.com/Juminiy/kube/pkg/internal/device"
)

func gpu() *sysGPU {
	return &sysGPU{Apple: internaldev.MPS()}
}
