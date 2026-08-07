package disk

import (
	"github.com/jianlu8023/go-tools/v2/pkg/bytes"
	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
	"github.com/shirou/gopsutil/v4/disk"
)

type Disk struct {
	MountPoint        string  `json:"mount_point,omitempty" yaml:"mount_point,omitempty"`
	Free              string  `json:"free,omitempty" yaml:"free,omitempty"`
	Used              string  `json:"used,omitempty" yaml:"used,omitempty"`
	Total             string  `json:"total,omitempty" yaml:"total,omitempty"`
	UsedPercentage    float64 `json:"used_percentage,omitempty" yaml:"used_percentage,omitempty"`
	InodesFree        uint64  `json:"inodes_free,omitempty" yaml:"inodes_free,omitempty"`
	InodesUsed        uint64  `json:"inodes_used,omitempty" yaml:"inodes_used,omitempty"`
	InodesTotal       uint64  `json:"inodes_total,omitempty" yaml:"inodes_total,omitempty"`
	InodesUsedPercent float64 `json:"inodes_used_percent,omitempty" yaml:"inodes_used_percent,omitempty"`
}

func (s Disk) String() string {
	str, _ := jsoniter.MarshalString(s)
	return str
}

// SystemDiskInfo 硬盘信息
// @function: SystemDiskInfo
// @description: 硬盘信息
// @return: d Disk, err error
func SystemDiskInfo() (d []Disk, err error) {
	return Disks(true)
}

func Disks(allPartition bool) (d []Disk, err error) {
	partitions, err := disk.Partitions(allPartition)
	if err != nil {
		return d, err
	}

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			// return d, err
			d = append(d, Disk{
				MountPoint: partition.Mountpoint,
			})
		} else {
			d = append(d, Disk{
				MountPoint:        partition.Mountpoint,
				Free:              bytes.HumanBinary(usage.Free),
				Used:              bytes.HumanBinary(usage.Used),
				Total:             bytes.HumanBinary(usage.Total),
				UsedPercentage:    usage.UsedPercent,
				InodesFree:        usage.InodesFree,
				InodesUsed:        usage.InodesUsed,
				InodesTotal:       usage.InodesTotal,
				InodesUsedPercent: usage.InodesUsedPercent,
			})
		}
		// fmt.Printf("当前挂在 %v 总容量 %v 已使用 %v 剩余 %v\n", partition.Mountpoint, usage.Total, usage.Used, usage.Free)
	}

	return d, nil
}
