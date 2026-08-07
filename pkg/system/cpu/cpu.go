package cpu

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type Cpu struct {
	Cpus  []float64 `json:"cpus,omitempty" yaml:"cpus,omitempty"`
	Cores int       `json:"cores,omitempty" yaml:"cores,omitempty"`
}

func (s Cpu) String() string {
	str, _ := jsoniter.MarshalString(s)
	return str
}

// SystemCpuInfo 获取CPU信息
// @function: SystemCpuInfo
// @description: CPU信息
// @return: c Cpu, err error
func SystemCpuInfo() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}
