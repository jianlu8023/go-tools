package all

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
	"github.com/jianlu8023/go-tools/v2/pkg/system/cpu"
	"github.com/jianlu8023/go-tools/v2/pkg/system/disk"
	"github.com/jianlu8023/go-tools/v2/pkg/system/os"
	"github.com/jianlu8023/go-tools/v2/pkg/system/ram"
)

type SystemInfo struct {
	Os   os.Os       `json:"os" yaml:"os"`
	Cpu  cpu.Cpu     `json:"cpu" yaml:"cpu"`
	Ram  ram.Ram     `json:"ram" yaml:"ram"`
	Disk []disk.Disk `json:"disk" yaml:"disk"`
}

func (s SystemInfo) String() string {
	marshalString, _ := jsoniter.MarshalString(s)
	return marshalString
}

func (s SystemInfo) GoString() string {
	return s.String()
}

func SystemInformation() (SystemInfo, error) {
	s := SystemInfo{}
	osInfo := os.SystemOsInfo()
	s.Os = osInfo
	cpuInfo, err := cpu.SystemCpuInfo()
	if err != nil {
		return s, err
	}
	s.Cpu = cpuInfo
	ramInfo, err := ram.SystemRamInfo()
	if err != nil {
		return s, err
	}
	s.Ram = ramInfo
	diskInfo, err := disk.SystemDiskInfo()
	if err != nil {
		return s, err
	}
	s.Disk = diskInfo
	return s, nil
}
