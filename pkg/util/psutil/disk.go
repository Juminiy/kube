package psutil

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/samber/lo"
	psdisk "github.com/shirou/gopsutil/v4/disk"
)

func disk() []psdisk.PartitionStat {
	diskParts, err := psdisk.Partitions(true)
	if err != nil {
		stdlog.ErrorF("psutil disk error: %s", err.Error())
		return nil
	}
	return diskParts
}

type DiskV2 struct {
	DiskUsageInfo []DiskUsageInfo
	Usages        []*psdisk.UsageStat `json:"disk_raw_usage"`
}

type DiskUsageInfo struct {
	Path  string
	Total string
	Free  string
	Used  string
}

func diskV2() DiskV2 {
	anyUsage := false
	usageList := lo.Map(disk(),
		func(item psdisk.PartitionStat, _ int) *psdisk.UsageStat {
			usageStat, err := psdisk.Usage(item.Mountpoint)
			if err != nil {
				anyUsage = true
				return usageStat
			}
			return nil
		})
	if !anyUsage {
		usageRoot, err := psdisk.Usage("/")
		if err == nil {
			usageList = []*psdisk.UsageStat{usageRoot}
		}
	}
	usageList = lo.Filter(usageList, func(item *psdisk.UsageStat, _ int) bool {
		return item != nil
	})

	return DiskV2{
		DiskUsageInfo: lo.Map(usageList, func(item *psdisk.UsageStat, _ int) (info DiskUsageInfo) {
			if item == nil {
				return info
			}
			info.Total = util.MeasureByte(int(item.Total))
			info.Free = util.MeasureByte(int(item.Free))
			info.Used = util.MeasureByte(int(item.Used))
			return info
		}),
		Usages: usageList}
}
