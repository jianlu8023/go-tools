package ram

import (
	"github.com/jianlu8023/go-tools/v2/pkg/bytes"
	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
	"github.com/shirou/gopsutil/v4/mem"
)

type Ram struct {
	Free           string  `json:"free,omitempty" yaml:"free,omitempty"`
	Available      string  `json:"available,omitempty" yaml:"available,omitempty"`
	Used           string  `json:"used,omitempty" yaml:"used,omitempty"`
	Total          string  `json:"total,omitempty" yaml:"total,omitempty"`
	UsedPercentage float64 `json:"used_percentage,omitempty" yaml:"used_percentage,omitempty"`
}

func (s Ram) String() string {
	str, _ := jsoniter.MarshalString(s)
	return str
}

// SystemRamInfo RAM信息
// @function: SystemRamInfo
// @description: RAM信息
// @return: r Ram, err error
func SystemRamInfo() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.Free = bytes.HumanBinary(u.Free)
		r.Available = bytes.HumanBinary(u.Available)
		r.Used = bytes.HumanBinary(u.Used)
		r.Total = bytes.HumanBinary(u.Total)
		r.UsedPercentage = u.UsedPercent
	}
	return r, nil
}
