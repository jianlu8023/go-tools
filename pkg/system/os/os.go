package os

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
	"runtime"
)

type Os struct {
	GOOS         string `json:"goos,omitempty" yaml:"goos,omitempty"`
	NumCPU       int    `json:"num_cpu,omitempty" yaml:"num_cpu,omitempty"`
	Compiler     string `json:"compiler,omitempty" yaml:"compiler,omitempty"`
	GoVersion    string `json:"go_version,omitempty" yaml:"go_version,omitempty"`
	NumGoroutine int    `json:"num_goroutine,omitempty" yaml:"num_goroutine,omitempty"`
}

func (o Os) String() string {
	marshalString, _ := jsoniter.MarshalString(o)
	return marshalString
}

func SystemOsInfo() Os {
	return Os{
		GOOS:         runtime.GOOS,
		NumCPU:       runtime.NumCPU(),
		Compiler:     runtime.Compiler,
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
	}
}
