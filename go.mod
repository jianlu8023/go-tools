module github.com/jianlu8023/go-tools/v2

go 1.22

require (
	// chainmaker.org/chainmaker/pb-go/v2 v2.3.0
	// chainmaker.org/chainmaker/sdk-go/v2 v2.3.0
	github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/go-resty/resty/v2 v2.13.1
	github.com/google/uuid v1.6.0
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.20
	github.com/shirou/gopsutil/v4 v4.24.10
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.23.0
)

replace (
	// chainmaker.org/chainmaker/pb-go/v2 => chainmaker.org/chainmaker/pb-go/v2 v2.3.0
	// chainmaker.org/chainmaker/sdk-go/v2 => chainmaker.org/chainmaker/sdk-go/v2 v2.3.0
	github.com/araddon/dateparse => github.com/araddon/dateparse v0.0.0-20210429162001-6b43995a97de
	github.com/json-iterator/go => github.com/json-iterator/go v1.1.12
	github.com/shirou/gopsutil/v4 => github.com/shirou/gopsutil/v4 v4.24.10
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.1 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
