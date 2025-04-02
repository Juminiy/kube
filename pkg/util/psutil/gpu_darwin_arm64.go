//go:build darwin && arm64

package psutil

func gpu() *sysGPU {
	return &sysGPU{Apple: nil}
}
